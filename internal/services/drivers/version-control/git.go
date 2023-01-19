package version_control

import (
	"fmt"
	"github.com/pkg/errors"
	"os"

	log "github.com/AppsFlyer/go-logger"
	"github.com/go-git/go-git/v5" /// with go modules disabled
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

const TerraformPath = "%s/%s"

type Git struct {
	log      log.Logger
	UserName string
	Token    string
}

func InitGitProvider(log log.Logger, userName, token string) *Git {
	return &Git{
		log:      log,
		UserName: userName,
		Token:    token,
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
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth:     &http.BasicAuth{Password: g.Token, Username: g.UserName},
	})

	return err
}
