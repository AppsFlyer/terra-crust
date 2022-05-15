package services

import (
	"fmt"
	"strings"
	"text/template"

	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/templates"
)

type TemplateApi struct {
	ApiFuncMap template.FuncMap
}

func NewTemplateApi() *TemplateApi {
	return &TemplateApi{
		ApiFuncMap: template.FuncMap{
			"SimpleWrap":        SimpleWrap,
			"ModuleDataWrapper": ModuleDataWrapper,
			"GetDefaults":       GetDefaults,
		},
	}
}

func ModuleDataWrapper(moduleName string, moduleData templates.ModuleData) map[string]interface{} {
	return map[string]interface{}{
		"ModuleName": moduleName,
		"ModuleData": moduleData,
	}
}

func SimpleWrap(moduleName string, moduleData map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"ModuleData":   moduleData,
		"VariableName": moduleName,
	}
}

func GetDefaults(moduleName string, modulesMap *templates.MainModuleTF) string {
	var sb strings.Builder
	for k := range modulesMap.Module[moduleName].SimpleLocals {
		sb.WriteString(fmt.Sprintf(main_default_var_row_template, k, moduleName, k))
	}

	for k := range modulesMap.Module[moduleName].MapLocals {
		sb.WriteString(fmt.Sprintf(main_default_var_row_template, k, moduleName, k))
	}

	return sb.String()
}
