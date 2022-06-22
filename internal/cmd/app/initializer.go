package app

import (
	logger "github.com/AppsFlyer/go-logger"
	"github.com/AppsFlyer/terra-crust/internal/services"
	"github.com/AppsFlyer/terra-crust/internal/services/drivers"
)

func InitTerraformGeneratorService(log logger.Logger) *services.Terraform {
	config := NewConfig(log)

	parserDriver := drivers.NewTerraformParser(log)
	parser := services.NewParser(log, parserDriver)
	templateHandler := services.NewTemplateHandler(log)
	tfSvc := services.NewTerraform(log, parser, templateHandler, config.GetString("LOCALS_TEMPLATE_PATH"), config.GetString("VARIABLE_TEMPLATE_PATH"), config.GetString("MAIN_TEMPLATE_PATH"))

	return tfSvc
}
