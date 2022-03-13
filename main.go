package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/internal/services/drivers"
)

// const tmpl = `variable {{ .ModuleName }} {
// 	description = {{.Description}}
// 	type = object({
// 			{{range $objName,$objType := .ObjectTypeMapping}}
// 			{{$objName}}=optional({{$objType}})
// 			{{end}}
// 	})
// 	default = {
// 			{{range $defaultKey,$defaultVal := .DefaultValues}}
// 			{{$defaultKey}}={{$defaultVal}}
// 			{{end}}
// 	}
//   }`

type outputExample struct {
	ModuleName        string
	Description       string
	ObjectTypeMapping map[string]string
	DefaultValues     map[string]string
}

//hclwrite.Tokens

func main() {

	parser := drivers.NewTerraformParser("./consul-sync")
	m, err := parser.Parse()
	if err != nil {
		fmt.Println("failed to parse", err.Error())
	}

	out := &outputExample{
		ModuleName:        "consul-sync",
		Description:       "test",
		ObjectTypeMapping: make(map[string]string, len(m.Variables)),
		DefaultValues:     make(map[string]string, len(m.Variables)),
	}

	for _, v := range m.Variables {
		if v.Default != nil {
			out.ObjectTypeMapping[v.Name] = strings.ReplaceAll(string(v.Type.Bytes()), " ", "")
			out.DefaultValues[v.Name] = string(v.Default.Bytes())
		}
	}

	// t := template.Must(template.New("tmpl").Parse(tmpl))

	// var tpl bytes.Buffer
	// if err := t.Execute(&tpl, out); err != nil {
	// 	fmt.Println("failed executer", err.Error())
	// }

	// parse the template
	tmpl, _ := template.ParseFiles("templates/module.tmpl")

	// create a new file
	file, _ := os.Create("final-module.tf")
	defer file.Close()

	// apply the template to the vars map and write the result to file.
	tmpl.Execute(file, out)

}
