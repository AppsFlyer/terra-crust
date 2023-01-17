// Copyright 2022 AppsFlyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/AppsFlyer/terra-crust/internal/services/templates"
)

type TemplateAPI struct {
	APIFuncMap *template.FuncMap
}

func NewTemplateAPI() *TemplateAPI {
	return &TemplateAPI{
		APIFuncMap: &template.FuncMap{
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
		sb.WriteString(fmt.Sprintf(mainDefaultVarRowTemplate, k, moduleName, k))
	}

	for k := range modulesMap.Module[moduleName].MapLocals {
		sb.WriteString(fmt.Sprintf(mainDefaultVarRowTemplate, k, moduleName, k))
	}

	return sb.String()
}
