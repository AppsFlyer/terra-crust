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

const main_default_var_row_template = "%s = local.%s.%s \n"

type Terraform struct {
	parser             *ModuleParser
	localsTemplatePath string
	objectTemplatePath string
	mainTemplatePath   string
}

func NewTerraform(parser *ModuleParser, localsTemplatePath, objectTemplatePath, mainTemplatePath string) *Terraform {
	return &Terraform{
		parser:             parser,
		localsTemplatePath: localsTemplatePath,
		objectTemplatePath: objectTemplatePath,
		mainTemplatePath:   mainTemplatePath,
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
			if v.Default != nil && string(v.Default.Bytes()) != `""` {
				out[k].ObjectTypeMapping[v.Name] = strings.ReplaceAll(string(v.Type.Bytes()), " ", "")
				out[k].DefaultValues[v.Name] = string(v.Default.Bytes())
			}
		}
	}

	return t.WriteTemplateToFile("module_variables.tf", t.objectTemplatePath, destinationPath, out)
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

		out.Module[k] = &templates.ModuleData{
			SimpleLocals: make(map[string]string),
			MapLocals:    make(map[string]templates.ComplexVariableData),
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

				separator := "="
				if strings.Contains(rawDefault, ":") {
					separator = ":"
				}

				for i := range splittedRawString {
					rawDataString := strings.Split(splittedRawString[i], separator)
					propertyName := strings.TrimSpace(rawDataString[0])

					propertyValue := strings.TrimSpace(strings.Join(rawDataString[1:], ":"))
					if propertyValue == `"[]"` {
						continue
					}

					//if property name is none string
					if !strings.Contains(propertyName, `"`) {
						propertyName = fmt.Sprintf(`"%s"`, propertyName)
					}

					if _, ok := out.Module[k].MapLocals[v.Name]; !ok {
						out.Module[k].MapLocals[v.Name] = make(templates.ComplexVariableData)
					}

					out.Module[k].MapLocals[v.Name][propertyName] = propertyValue

					continue
				}
			}
		}
	}

	return t.WriteTemplateToFile("module_locals.tf", t.localsTemplatePath, destinationPath, out)
}

// func (t *Terraform) getVariableDetails(rawVariableData string) (string, string, error) {

// }

func (t *Terraform) GenerateMain(modulesFilePath, destinationPath, mainTemplatePath string) error {
	moduleList, err := t.parser.GetModulesList(modulesFilePath)
	if err != nil {
		return err
	}

	out := &templates.MainModuleTF{
		Module:   make(map[string]*templates.MainModuleData),
		RootPath: modulesFilePath,
	}

	for k, m := range moduleList {
		if len(m.Variables) == 0 {
			continue
		}

		out.Module[k] = &templates.MainModuleData{
			ModuleData: &templates.ModuleData{
				SimpleLocals: make(map[string]string),
				MapLocals:    make(map[string]templates.ComplexVariableData),
			},
			RequiredFields: make(map[string]string),
		}

		for _, v := range m.Variables {
			//Simple variable
			if v.Default != nil && string(v.Default.Bytes()) != `""` && !strings.Contains(string(v.Type.Bytes()), "map") {
				out.Module[k].SimpleLocals[v.Name] = string(v.Default.Bytes())
			}

			//Map variable
			if v.Default != nil && string(v.Default.Bytes()) != `""` && strings.Contains(string(v.Type.Bytes()), "map") {
				rawDefault := string(v.Default.Bytes())
				rawDefault = strings.ReplaceAll(rawDefault, "{", "")
				rawDefault = strings.ReplaceAll(rawDefault, "}", "")
				rawDefault = strings.TrimSpace(rawDefault)

				splittedRawString := strings.Split(rawDefault, "\n")

				separator := "="
				if strings.Contains(rawDefault, ":") {
					separator = ":"
				}
				for i := range splittedRawString {
					rawDataString := strings.Split(splittedRawString[i], separator)
					propertyName := strings.TrimSpace(rawDataString[0])
					if _, ok := out.Module[k].MapLocals[v.Name]; !ok {
						out.Module[k].MapLocals[v.Name] = make(templates.ComplexVariableData)
					}

					out.Module[k].MapLocals[v.Name][propertyName] = ""
				}
			}

			//Required Variable
			if v.Default == nil || string(v.Default.Bytes()) == `""` {
				out.Module[k].RequiredFields[v.Name] = ""
			}
		}
	}
	path := t.mainTemplatePath
	if mainTemplatePath != "" {
		path = mainTemplatePath
	}

	return t.WriteTemplateToFile("module_main.tf", path, destinationPath, out)
}

func (t *Terraform) WriteTemplateToFile(fileName, templatePath, destinationPath string, out interface{}) error {
	splittedPath := strings.Split(templatePath, "/")
	templateName := splittedPath[len(splittedPath)-1]
	tmpl, err := template.New(templateName).Funcs(NewTemplateApi().ApiFuncMap).ParseFiles(templatePath)
	if err != nil {
		//fmt.Println(err.Error()) TODO: add logger
		return err
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, out); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", destinationPath, fileName)
	if err := os.Remove(filePath); (err != nil) && (!errors.Is(err, os.ErrNotExist)) {
		return err
	}

	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	_, err = file.WriteString(buf.String())
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
