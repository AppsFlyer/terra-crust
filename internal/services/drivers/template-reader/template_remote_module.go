package template_reader

import (
	"errors"
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
)

type KeyValueString struct {
	Key   string
	Value string
}

type TemplateRemoteModule struct {
	log     log.Logger
	modules map[string]string
}

func InitTemplateRemoteModule(log log.Logger) *TemplateRemoteModule {
	return &TemplateRemoteModule{
		log:     log,
		modules: make(map[string]string),
	}
}

func (trm *TemplateRemoteModule) GetRemoteModulesFromCache() map[string]string {
	return trm.modules
}

func (trm *TemplateRemoteModule) GetRemoteModulesFromTemplate(templatePath string) (map[string]string, error) {
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

		trm.modules[kv.Key] = kv.Value
	}

	return trm.modules, nil
}

func (trm *TemplateRemoteModule) parseLineIntoRemoteModule(line string) (*KeyValueString, error) {
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

	if len(urlMatch) >= 1 {
		gitUrl := strings.TrimSuffix(urlMatch[1], "\"")
		return &KeyValueString{Key: moduleName, Value: gitUrl}, nil
	}
	gitUrl := strings.TrimSuffix(urlVersionMatch[1], "\"")

	return &KeyValueString{Key: moduleName, Value: gitUrl}, nil
}
