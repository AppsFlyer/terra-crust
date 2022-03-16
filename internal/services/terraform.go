package services

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/templates"
)

const moduleDescription = `<<EOT
	(Optional) %s Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/%s/README.md)
	EOT`

type Terraform struct {
	parser             *ModuleParser
	localsTemplatePath string
	objectTemplatePath string
}

func NewTerraform(parser *ModuleParser, localsTemplatePath, objectTemplatePath string) *Terraform {
	return &Terraform{
		parser:             parser,
		localsTemplatePath: localsTemplatePath,
		objectTemplatePath: objectTemplatePath,
	}
}

func (t *Terraform) GenerateModuleVariableObject(modulesFilePath, destinationPath string) error {
	moduleList, err := t.parser.GetModulesList(modulesFilePath)
	if err != nil {
		return err
	}

	out := make(templates.VariblesModuleList)

	for k, m := range moduleList {
		if len(m.Variables) == 0 {
			continue
		}

		out[k] = &templates.VariablesModulesTF{
			ModuleName:        k,
			Description:       fmt.Sprintf(moduleDescription, k, k),
			ObjectTypeMapping: make(map[string]string, len(m.Variables)),
			DefaultValues:     make(map[string]string, len(m.Variables)),
		}

		for _, v := range m.Variables {
			fmt.Printf("module name :%s var name: %s default: %s \n", k, v.Name, string(v.Default.Bytes()))
			if v.Default != nil && string(v.Default.Bytes()) != `""` {
				out[k].ObjectTypeMapping[v.Name] = strings.ReplaceAll(string(v.Type.Bytes()), " ", "")
				out[k].DefaultValues[v.Name] = string(v.Default.Bytes())
			}
		}
	}

	return t.WriteTemplateToFile(t.objectTemplatePath, destinationPath, out)
}

func (t *Terraform) GenerateModuleDefaultLocals(modulesFilePath, destinationPath string) error {
	moduleList, err := t.parser.GetModulesList(modulesFilePath)
	if err != nil {
		return err
	}

	out := &templates.LocalsModulesTF{
		Module: make(map[string]*templates.ModuleData),
	}

	for k, m := range moduleList {
		if len(m.Variables) == 0 {
			continue
		}
		fmt.Printf("module name :%s \n", k)

		out.Module[k] = &templates.ModuleData{
			SimpleLocals:    make(map[string]string),
			ComplexLocals:   make(map[string]templates.ComplexVariableData),
			LocalsStringKey: make(map[string]templates.ComplexVariableData),
		}

		for _, v := range m.Variables {
			if v.Default != nil && string(v.Default.Bytes()) != `""` && !strings.Contains(string(v.Type.Bytes()), "map") {
				out.Module[k].SimpleLocals[v.Name] = string(v.Default.Bytes())
			}

			if v.Default != nil && string(v.Default.Bytes()) != `""` && strings.Contains(string(v.Type.Bytes()), "map") {
				rawDefault := string(v.Default.Bytes())
				rawDefault = strings.ReplaceAll(rawDefault, "{", "")
				rawDefault = strings.ReplaceAll(rawDefault, "}", "")
				rawDefault = strings.TrimSpace(rawDefault)

				splittedRawString := strings.Split(rawDefault, "\n")

				seperator := "="
				if strings.Contains(rawDefault, ":") {
					seperator = ":"
				}

				for i := range splittedRawString {
					rawDataString := strings.Split(splittedRawString[i], seperator)
					propertyName := strings.TrimSpace(rawDataString[0])

					propertyValue := strings.TrimSpace(strings.Join(rawDataString[1:], ":"))
					if propertyValue == `"[]"` {
						continue
					}

					if strings.Contains(propertyName, `"`) {
						if _, ok := out.Module[k].LocalsStringKey[v.Name]; !ok {
							out.Module[k].LocalsStringKey[v.Name] = make(templates.ComplexVariableData)
						}

						out.Module[k].LocalsStringKey[v.Name][propertyName] = propertyValue
						continue
					}

					if _, ok := out.Module[k].ComplexLocals[v.Name]; !ok {
						out.Module[k].ComplexLocals[v.Name] = make(templates.ComplexVariableData)
					}

					out.Module[k].ComplexLocals[v.Name][propertyName] = propertyValue
				}
			}
		}
	}

	return t.WriteTemplateToFile(t.localsTemplatePath, destinationPath, out)
}

func (t *Terraform) WriteTemplateToFile(templatePath, destinationPath string, out interface{}) error {
	tmpl := template.Must(template.ParseFiles(templatePath))
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, out); err != nil {
		return err
	}

	if err := os.Remove(destinationPath); (err != nil) && (!errors.Is(err, os.ErrNotExist)) {
		return err
	}

	file, _ := os.OpenFile(destinationPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	_, err := file.WriteString(buf.String())
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
