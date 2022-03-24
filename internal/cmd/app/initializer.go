package app

import (
	logger "gitlab.appsflyer.com/go/af-go-logger/v1"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
)

func InitTerraformGeneratorService(log logger.Logger) *services.Terraform {
	config := NewConfig(log)

	parserDriver := drivers.NewTerraformParser()
	parser := services.NewParser(parserDriver)
	tfSvc := services.NewTerraform(parser, config.GetString("LOCALS_TEMPLATE_PATH"), config.GetString("VARIABLE_TEMPLATE_PATH"))

	return tfSvc
}
