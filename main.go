package main

import (
	"fmt"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
)

func main() {

	driver := drivers.NewTerraformParser()
	parser := services.NewParser(driver, "./internal/templates/variables_modules.tf.tmpl")

	terraform := services.NewTerraform(parser, "./internal/templates/locals_modules.tf.tmpl", "./internal/templates/variables_modules.tf.tmpl")

	if err := terraform.GenerateModuleVariableObject("./modules", "final-module.tf"); err != nil {
		fmt.Println(err.Error())
	}

	err := terraform.GenerateModuleDefaultLocals("./modules", "test-module.tf")
	if err != nil {
		fmt.Println(err.Error())
	}
}
