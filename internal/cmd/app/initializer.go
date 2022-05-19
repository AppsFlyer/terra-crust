package app

import (
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/services/drivers"
)

func InitTerraformGeneratorService(log logger.Logger) *services.Terraform {
	config := NewConfig(log)

	parserDriver := drivers.NewTerraformParser()
	parser := services.NewParser(parserDriver)
	templateHandler := services.NewTemplateHandler(log)
	tfSvc := services.NewTerraform(parser, templateHandler, config.GetString("LOCALS_TEMPLATE_PATH"), config.GetString("VARIABLE_TEMPLATE_PATH"), config.GetString("MAIN_TEMPLATE_PATH"))

	return tfSvc
}
