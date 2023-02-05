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

package app

import (
	logger "github.com/AppsFlyer/go-logger"
	"github.com/AppsFlyer/terra-crust/internal/services"
	"github.com/AppsFlyer/terra-crust/internal/services/drivers/parser"
)

const (
	LocalsTemplatePath   = "templates/locals_file.tmpl"
	VariableTemplatePath = "templates/variables_file.tmpl"
	MainTemplatePath     = "templates/main_file.tmpl"
)

func InitTerraformGeneratorService(log logger.Logger) *services.Terraform {
	parserDriver := parser.NewTerraformParser(log)
	parserSvc := services.NewParser(log, parserDriver)
	templateHandler := services.NewTemplateHandler(log)

	tfSvc := services.NewTerraform(log, parserSvc, templateHandler, LocalsTemplatePath, VariableTemplatePath, MainTemplatePath)

	return tfSvc
}
