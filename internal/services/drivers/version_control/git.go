package version_control

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	log "github.com/AppsFlyer/go-logger"
	"github.com/go-git/go-git/v5" /// with go modules disabled
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	cp "github.com/otiai10/copy"
	"github.com/pkg/errors"
)

const (
	FolderPathFormat            = "%s/%s"
	TempFolderPath              = "%s/temp_clone_path/%s"
	remoteName                  = "origin"
	GitlabTokenENV              = "GITLAB_TOKEN"
	GitRefTag                   = "refs/tags/%s"
	GitRefBranch                = "refs/remotes/%s/%s"
	GitCredentialUrl            = "url=%s"
	GitCredentialDeadLineMs     = 500
	GitCredentialUserNamePrefix = "username="
	GitCredentialPasswordPrefix = "password="
	GitlabUserENV               = "GITLAB_USER"
	GithubTokenENV              = "GITHUB_TOKEN"
	GithubUserENV               = "GITHUB_USER"
)

type RemoteModule struct {
	Name    string
	Url     string
	Version string
	Path    string
}

type Git struct {
	log log.Logger
}

func InitGitProvider(log log.Logger) *Git {
	return &Git{
		log: log,
	}
}

func (g *Git) CloneModules(modules map[string]*RemoteModule, modulesSource string, externalGit bool) error {
	for moduleName, moduleData := range modules {
		clonePath := fmt.Sprintf(FolderPathFormat, modulesSource, moduleName)
		if moduleData.Path != "" {
			clonePath = fmt.Sprintf(TempFolderPath, modulesSource, moduleName)
		}

		if err := g.clone(moduleData, clonePath, externalGit); err != nil {
			return errors.Wrapf(err, "failed to fetch module , url: %s", moduleData.Url)
		}

		g.log.Infof("Copy folder from : %s/%s to : %s", clonePath, moduleData.Path, modulesSource)

		if moduleData.Path != "" {
			modulePath := fmt.Sprintf(FolderPathFormat, modulesSource, moduleName)

			err := cp.Copy(fmt.Sprintf(FolderPathFormat, clonePath, moduleData.Path), modulePath)
			if err != nil {
				g.log.Errorf("failed to copy desired terraform module module path :%s, module name: %s", clonePath, moduleData.Name)
			}

			g.log.Infof("Changing module path from : %s to : %s", modules[moduleName].Path, modulePath)
			modules[moduleName].Path = moduleName
		}
	}

	if err := g.cleanTemp(modulesSource); err != nil {
		return errors.Wrap(err, "Failed to delete temp folder")
	}

	return nil
}

func (g *Git) clone(moduleData *RemoteModule, directoryPath string, externalGit bool) error {

	userName, token, err := g.getGitCredentials(moduleData.Url, externalGit)
	if err != nil {
		return err
	}

	repo, err := git.PlainClone(directoryPath, false, &git.CloneOptions{
		URL:        moduleData.Url,
		Auth:       &http.BasicAuth{Password: token, Username: userName},
		RemoteName: remoteName,
		Depth:      1,
	})
	if err != nil {
		return err
	}
	if moduleData.Version != "" {
		workTree, err := repo.Worktree()
		if err != nil {
			return err
		}
		tagRef := fmt.Sprintf(GitRefTag, moduleData.Version)
		g.log.Debugf("searching %s in %s", moduleData.Version, tagRef)
		err = workTree.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName(tagRef),
		})
		if err != nil {
			branchRef := fmt.Sprintf(GitRefBranch, remoteName, moduleData.Version)
			g.log.Debugf("version not found in tags ref, searching %s in %s", moduleData.Version, branchRef)
			bErr := workTree.Checkout(&git.CheckoutOptions{
				Branch: plumbing.ReferenceName(branchRef),
			})
			return bErr
		}
	}
	return nil
}

func (g *Git) CleanModulesFolders(modules map[string]*RemoteModule, modulesSource string) error {
	var returnedErr error = nil
	for moduleName := range modules {
		modulePath := fmt.Sprintf(FolderPathFormat, modulesSource, moduleName)

		err := os.RemoveAll(modulePath)
		if err != nil {
			g.log.Errorf("Failed to clear up a module %s at path: %s , error:%s", moduleName, modulePath, err.Error())
			if returnedErr == nil {
				returnedErr = err
			}

			returnedErr = errors.Wrap(returnedErr, err.Error())
		}
	}

	return returnedErr
}

func (g *Git) cleanTemp(modulesSourcePath string) error {
	tempPath := fmt.Sprintf("%s/temp_clone_path", modulesSourcePath)

	g.log.Infof("Deleting temp folder containing all clones")

	err := os.RemoveAll(tempPath)
	if err != nil {
		g.log.Errorf("Failed to clear up temp folder that been used for clone, please clean it manually , failed on temp folder, path : %s , err: %s", tempPath, err.Error())

		return err
	}

	return nil
}

func (g *Git) getGitUserNameAndToken(url string) (string, string) {
	if strings.Contains(url, "gitlab") {
		return os.Getenv(GitlabUserENV), os.Getenv(GitlabTokenENV)
	}

	if strings.Contains(url, "github") {
		return os.Getenv(GithubUserENV), os.Getenv(GithubTokenENV)
	}

	return "", ""
}
func (g *Git) getGitCredentials(url string, externalGit bool) (userName string, password string, err error) {
	if !externalGit {
		userName, password = g.getGitUserNameAndToken(url)
		return userName, password, nil
	}
	// Required until https://github.com/go-git/go-git/issues/490 addressed
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(GitCredentialDeadLineMs*time.Millisecond))
	cmd := exec.CommandContext(ctx, "git", "credential", "fill")
	defer cancel()
	cmd.Stdin = strings.NewReader(fmt.Sprintf(GitCredentialUrl, url))
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return userName, password, err
	}
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, GitCredentialUserNamePrefix) {
			userName = strings.TrimPrefix(line, GitCredentialUserNamePrefix)
		}
		if strings.HasPrefix(line, GitCredentialPasswordPrefix) {
			password = strings.TrimPrefix(line, GitCredentialPasswordPrefix)
		}
	}
	g.log.Infof("user=%s,password=%s", userName, password)
	return userName, password, nil
}
