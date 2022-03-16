package services

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/templates"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/types"
)

type ModuleParser struct {
	parsingDriver drivers.Parser
	templatePath  string
}

const moduleDescription = `<<EOT
	(Optional) %s Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/%s/README.md)
	EOT`

func NewParser(driver drivers.Parser, templatePath string) *ModuleParser {
	return &ModuleParser{
		parsingDriver: driver,
		templatePath:  templatePath,
	}
}

func (p *ModuleParser) ParseModulesVariables(rootFolder string) (*bytes.Buffer, error) {
	modulesList, err := p.parsingDriver.Parse(rootFolder)
	if err != nil {
		fmt.Println("failed to parse", err.Error())

		return nil, err
	}

	buf := new(bytes.Buffer)

	for k, m := range modulesList {
		if len(m.Variables) == 0 {
			continue
		}

		out := &templates.VariablesModulesTF{
			ModuleName:        k,
			Description:       fmt.Sprintf(moduleDescription, k, k),
			ObjectTypeMapping: make(map[string]string, len(m.Variables)),
			DefaultValues:     make(map[string]string, len(m.Variables)),
		}

		for _, v := range m.Variables {
			if v.Default != nil && string(v.Default.Bytes()) != `""` {
				out.ObjectTypeMapping[v.Name] = strings.ReplaceAll(string(v.Type.Bytes()), " ", "")
				out.DefaultValues[v.Name] = string(v.Default.Bytes())
			}
		}

		tmpl := template.Must(template.ParseFiles(p.templatePath))
		err = tmpl.Execute(buf, out)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func (p *ModuleParser) GetModulesList(rootFolder string) (map[string]*types.Module, error) {
	modulesList, err := p.parsingDriver.Parse(rootFolder)
	if err != nil {
		fmt.Println("failed to parse", err.Error())

		return nil, err
	}

	return modulesList, nil
}
