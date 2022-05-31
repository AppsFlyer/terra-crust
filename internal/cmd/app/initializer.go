package app

import (
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services/drivers"
)

func InitTerraformGeneratorService(log logger.Logger) *services.Terraform {
	config := NewConfig(log)

	parserDriver := drivers.NewTerraformParser(log)
	parser := services.NewParser(log, parserDriver)
	templateHandler := services.NewTemplateHandler(log)
	tfSvc := services.NewTerraform(log, parser, templateHandler, config.GetString("LOCALS_TEMPLATE_PATH"), config.GetString("VARIABLE_TEMPLATE_PATH"), config.GetString("MAIN_TEMPLATE_PATH"))

	return tfSvc
}
