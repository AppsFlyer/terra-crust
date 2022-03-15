package main

import (
	"fmt"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
)

// func main() {

// 	parser := drivers.NewTerraformParser("./")
// 	m, err := parser.Parse()
// 	if err != nil {
// 		fmt.Println("failed to parse", err.Error())
// 	}

// 	// for _, v := range m.Variables {
// 	// 	fmt.Println("name:")
// 	// 	fmt.Println(v.Name)
// 	// 	fmt.Println("default:")
// 	// 	fmt.Println(strings.ReplaceAll(string(v.Default.Bytes()), " ", ""))
// 	// }

// 	fmt.Println(m)

// }

//hclwrite.Tokens

func main() {

	driver := drivers.NewTerraformParser()
	parser := services.NewParser(driver, "./internal/templates/variables_modules.tf.tmpl")

	terraform := services.NewTerraform(parser, "./internal/templates/locals_modules.tf.tmpl")

	// if err := terraform.GenerateSubModule("./modules", "final-module.tf"); err != nil {
	// 	fmt.Println(err.Error())
	// }

	err := terraform.GenerateModuleLocals("./modules", "test-module.tf")
	if err != nil {
		fmt.Println(err.Error())
	}

	// create-sub-module path
	// v2 multiple paths
	// override/path
}
