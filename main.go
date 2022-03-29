package main

import (
	"fmt"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
)

func main() {

	driver := drivers.NewTerraformParser()
	parser := services.NewParser(driver)

	terraform := services.NewTerraform(parser, "./internal/templates/locals_file.tmpl", "./internal/templates/variables_file.tmpl", "./internal/templates/main_file.tmpl")

	// if err := terraform.GenerateModuleVariableObject("./modules", "."); err != nil {
	// 	fmt.Println(err.Error())
	// }

	err := terraform.GenerateModuleDefaultLocals("./modules", ".")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := terraform.GenerateMain("./modules", "."); err != nil {
		fmt.Println(err.Error())
	}
}
