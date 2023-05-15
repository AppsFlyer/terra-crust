package version_control

import (
	"fmt"
	"os"
	"strings"

	log "github.com/AppsFlyer/go-logger"
	"github.com/go-git/go-git/v5" /// with go modules disabled
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	cp "github.com/otiai10/copy"
	"github.com/pkg/errors"
)

const (
	FolderPathFormat = "%s/%s"
	TempFolderPath   = "%s/temp_clone_path/%s"
	GitlabTokenENV   = "GITLAB_TOKEN"
	GitlabUserENV    = "GITLAB_USER"
	GithubTokenENV   = "GITHUB_TOKEN"
	GithubUserENV    = "GITHUB_USER"
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

func (g *Git) CloneModules(modules map[string]*RemoteModule, modulesSource string) error {
	for moduleName, moduleData := range modules {
		clonePath := fmt.Sprintf(FolderPathFormat, modulesSource, moduleName)
		if moduleData.Path != "" {
			clonePath = fmt.Sprintf(TempFolderPath, modulesSource, moduleName)
		}

		if err := g.clone(moduleData, clonePath); err != nil {
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

func (g *Git) clone(moduleData *RemoteModule, directoryPath string) error {
	userName, token := g.getGitUserNameAndToken(moduleData.Url)

	remoteName := "origin"
	if moduleData.Version != "" {
		remoteName = moduleData.Version
	}

	_, err := git.PlainClone(directoryPath, false, &git.CloneOptions{
		URL:        moduleData.Url,
		Auth:       &http.BasicAuth{Password: token, Username: userName},
		RemoteName: remoteName,
		Depth:      1,
	})

	return err
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
