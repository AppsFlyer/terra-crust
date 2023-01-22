package version_control

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"

	log "github.com/AppsFlyer/go-logger"
	"github.com/go-git/go-git/v5" /// with go modules disabled
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const (
	TerraformPath  = "%s/%s"
	GitlabTokenENV = "GITLAB_TOKEN"
	GithubTokenENV = "GITHUB_TOKEN"
	GitlabUserENV  = "GITLAB_USER"
	GithubUserENV  = "GITHUB_USER"
)

type Git struct {
	log log.Logger
}

func InitGitProvider(log log.Logger) *Git {
	return &Git{
		log: log,
	}
}

func (g *Git) CloneModules(modules map[string]string, modulesSource string) error {
	for moduleName, moduleUrl := range modules {
		if err := g.clone(moduleUrl, fmt.Sprintf(TerraformPath, modulesSource, moduleName)); err != nil {
			return err
		}
	}

	return nil
}

func (g *Git) CleanModulesFolders(modules map[string]string, modulesSource string) error {
	var returnedErr error = nil
	for moduleName := range modules {
		modulePath := fmt.Sprintf(TerraformPath, modulesSource, moduleName)
		err := os.RemoveAll(modulePath)
		if err != nil {
			g.log.Errorf("Failed to clearup a module %s at path: %s , error:%s", moduleName, modulePath, err.Error())
			if returnedErr == nil {
				returnedErr = err
			}

			returnedErr = errors.Wrap(returnedErr, err.Error())
		}
	}

	return returnedErr
}

func (g *Git) clone(url, path string) error {
	userName, token := g.getGitUserNameAndToken(url)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth:     &http.BasicAuth{Password: token, Username: userName},
	})

	return err
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
