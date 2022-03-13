package drivers

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
	"gitlab.appsflyer.com/real-time-platform/terraform-submodule-wrapper/types"
)

type TerraformParser struct {
	Source string
}

// newLocalParser return a new parser with local terraform module.
func NewTerraformParser(source string) Parser {
	return &TerraformParser{
		Source: source,
	}
}

// Parse parses a local terraform module and returns module structs
func (p *TerraformParser) Parse() (*types.Module, error) {
	if !(strings.HasPrefix(p.Source, "./") || strings.HasPrefix(p.Source, "../")) {
		return nil, errors.New("Invalid local module path.")
	}

	var variables []*types.Variable
	var outputs []*types.Output
	var resources []*types.Resource

	err := filepath.Walk(p.Source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() || !strings.HasSuffix(info.Name(), ".tf") {
				return nil
			}

			src, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			file, diags := hclwrite.ParseConfig(src, path, hcl.InitialPos)
			if diags.HasErrors() {
				return diags
			}

			body := file.Body()
			for _, block := range body.Blocks() {
				switch block.Type() {
				case "variable":
					variables = append(variables, p.parseVariable(block))
				case "output":
					outputs = append(outputs, p.parseOutput(block))
				case "resource":
					resources = append(resources, p.parseResource(block))
				}
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	return &types.Module{
		Source:    p.Source,
		Variables: variables,
		Outputs:   outputs,
		Resources: resources,
	}, nil
}

func (p *TerraformParser) parseVariable(block *hclwrite.Block) *types.Variable {
	variable := types.Variable{
		Name:    block.Labels()[0],
		Default: hclwrite.TokensForValue(cty.StringVal("")),
	}
	body := block.Body()
	for k, v := range body.Attributes() {
		switch k {
		case "type":
			var typeTokens hclwrite.Tokens
			for _, t := range v.Expr().BuildTokens(nil) {
				if t.Type != hclsyntax.TokenNewline {
					typeTokens = append(typeTokens, t)
				}
			}
			variable.Type = typeTokens
		case "default":
			variable.Default = v.Expr().BuildTokens(nil)
		case "description":
			description := string(v.Expr().BuildTokens(nil).Bytes())
			variable.Description = description[2 : len(description)-1]
		}
	}
	return &variable
}

func (p *TerraformParser) parseOutput(block *hclwrite.Block) *types.Output {
	output := types.Output{
		Name:        block.Labels()[0],
		Description: "",
	}
	body := block.Body()
	for k, v := range body.Attributes() {
		switch k {
		case "value":
			var typeTokens hclwrite.Tokens
			for _, t := range v.Expr().BuildTokens(nil) {
				typeTokens = append(typeTokens, t)
			}
			output.Value = typeTokens
		case "description":
			description := string(v.Expr().BuildTokens(nil).Bytes())
			output.Description = description[2 : len(description)-1]
		}
	}
	return &output
}

func (p *TerraformParser) parseResource(block *hclwrite.Block) *types.Resource {
	resource := types.Resource{
		Type: block.Labels()[0],
		Name: block.Labels()[1],
	}
	return &resource
}
