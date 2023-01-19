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
	gitSourceLineRegex = "source\\s*=\\s*git::"
	moduleNameRegex    = `\/([^\/]+)\.git(?:[^\/]*|$)`
	moduleUrlRegex     = "git::(https?:\\/\\/[^\\/]+\\/[^?]+.git)"
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
	defer file.Close()

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
		return nil, errors.New("Failed to fetch module Name from source")
	}
	moduleName := match[1]

	re = regexp.MustCompile(moduleUrlRegex)
	match = re.FindStringSubmatch(line)
	if len(match) <= 1 {
		return nil, errors.New("Failed to fetch module URL from source")
	}
	gitUrl := strings.TrimPrefix(match[1], "git::")
	return &KeyValueString{Key: moduleName, Value: gitUrl}, nil
}
