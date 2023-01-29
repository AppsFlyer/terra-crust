package template_reader

import (
	"errors"
	version_control "github.com/AppsFlyer/terra-crust/internal/services/drivers/version-control"
	"io"
	"os"
	"regexp"
	"strings"

	log "github.com/AppsFlyer/go-logger"
)

const (
	gitSourceLineRegex = "source\\s*=\\s*\"git::"
	moduleNameRegex    = `.*\/([^?]+)`
	gitUrlRegex        = `(https?:\/\/.+\/.+\.git)`
	gitUrlRegexVersion = `(https?:\/\/[^/]+\/[^?]+)`
	versionRegex       = `/?ref=(.*)"`
	modulePathRegex    = `.git[\/]+([^\/][^\?]+)[\?]{0,1}.*["$]`
)

type TemplateRemoteModule struct {
	log     log.Logger
	modules map[string]*version_control.RemoteModule
}

func InitTemplateRemoteModule(log log.Logger) *TemplateRemoteModule {
	return &TemplateRemoteModule{
		log:     log,
		modules: make(map[string]*version_control.RemoteModule),
	}
}

func (trm *TemplateRemoteModule) GetRemoteModulesFromCache() map[string]*version_control.RemoteModule {
	return trm.modules
}

func (trm *TemplateRemoteModule) GetRemoteModulesFromTemplate(templatePath string) (map[string]*version_control.RemoteModule, error) {
	file, err := os.Open(templatePath)
	if err != nil {
		trm.log.Error("Failed to open template to fetch remote modules,make sure you have a custom template when using this feature")

		return nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			trm.log.Errorf("Failed to close properly the main tmpl , error: %s", err.Error())
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		trm.log.Error("Failed to read file")

		return nil, err
	}

	r, err := regexp.Compile(gitSourceLineRegex)
	if err != nil {
		trm.log.Error("Failed to compile regex")

		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if !r.MatchString(line) {
			continue
		}

		kv, err := trm.parseLineIntoRemoteModule(line)
		if err != nil {
			trm.log.Error("Failed to parse remote Source value from template")

			return nil, err
		}

		trm.modules[kv.Name] = kv
	}

	return trm.modules, nil
}

func (trm *TemplateRemoteModule) parseLineIntoRemoteModule(line string) (*version_control.RemoteModule, error) {
	re := regexp.MustCompile(moduleNameRegex)
	match := re.FindStringSubmatch(line)
	if len(match) <= 1 {
		return nil, errors.New("failed to fetch module Name from source")
	}
	moduleName := strings.TrimSuffix(match[1], "\"")

	urlReg := regexp.MustCompile(gitUrlRegex)
	urlVersionReg := regexp.MustCompile(gitUrlRegexVersion)
	urlMatch := urlReg.FindStringSubmatch(line)
	urlVersionMatch := urlVersionReg.FindStringSubmatch(line)

	if len(urlMatch) < 1 && len(urlVersionMatch) < 1 {
		return nil, errors.New("failed to fetch module URL from source")
	}

	versReg := regexp.MustCompile(versionRegex)
	verMatch := versReg.FindStringSubmatch(line)
	version := ""
	if len(verMatch) >= 1 {
		version = verMatch[1]
	}

	modulePathReg := regexp.MustCompile(modulePathRegex)
	modulePathMatch := modulePathReg.FindStringSubmatch(line)
	modulePath := ""
	if len(modulePathMatch) >= 1 {
		modulePath = modulePathMatch[1]
	}

	if len(urlMatch) >= 1 {
		gitUrl := strings.TrimSuffix(urlMatch[1], "\"")

		return &version_control.RemoteModule{Name: moduleName, Url: gitUrl, Version: version, Path: modulePath}, nil
	}

	gitUrl := strings.TrimSuffix(urlVersionMatch[1], "\"")

	return &version_control.RemoteModule{Name: moduleName, Url: gitUrl, Version: version, Path: modulePath}, nil
}
