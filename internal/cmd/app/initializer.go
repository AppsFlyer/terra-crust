package app

import (
	logger "github.com/AppsFlyer/go-logger"
	"github.com/AppsFlyer/terra-crust/internal/services"
	"github.com/AppsFlyer/terra-crust/internal/services/drivers"
)

const LOCALS_TEMPLATE_PATH = "templates/locals_file.tmpl"
const VARIABLE_TEMPLATE_PATH = "templates/variables_file.tmpl"
const MAIN_TEMPLATE_PATH = "templates/main_file.tmpl"

func InitTerraformGeneratorService(log logger.Logger) *services.Terraform {
	parserDriver := drivers.NewTerraformParser(log)
	parser := services.NewParser(log, parserDriver)
	templateHandler := services.NewTemplateHandler(log)

	tfSvc := services.NewTerraform(log, parser, templateHandler, LOCALS_TEMPLATE_PATH, VARIABLE_TEMPLATE_PATH, MAIN_TEMPLATE_PATH)

	return tfSvc
}
