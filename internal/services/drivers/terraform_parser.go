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
	log "gitlab.appsflyer.com/go/af-go-logger/v1"
	"gitlab.appsflyer.com/real-time-platform/terra-crust/internal/types"
)

type TerraformParser struct {
	logger log.Logger
}

// newLocalParser return a new parser with local terraform module.
func NewTerraformParser(logger log.Logger) Parser {
	return &TerraformParser{
		logger: logger,
	}
}

// Parse parses a local terraform module and returns module structs
func (p *TerraformParser) Parse(path string) (map[string]*types.Module, error) {
	modulesMap := make(map[string]*types.Module)
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				p.logger.ErrorWithError("failed scraping folder on filepath.walk", err)

				return err
			}

			if info.IsDir() || !strings.HasSuffix(info.Name(), ".tf") {
				p.logger.ErrorWithError("Folder is empty or there is no files with .tf please make sure path is correct", err)

				return nil
			}

			src, err := ioutil.ReadFile(path)
			if err != nil {
				p.logger.ErrorWithError("Failed reading file under path, make sure files exist", err)

				return err
			}

			file, diags := hclwrite.ParseConfig(src, path, hcl.InitialPos)
			if diags.HasErrors() {
				return errors.New(diags.Error())
			}

			body := file.Body()
			splittedPath := strings.Split(path, "/")
			//Extracting module name based on folder file location
			moduleName := splittedPath[len(splittedPath)-2]

			if _, ok := modulesMap[moduleName]; !ok {
				modulesMap[moduleName] = &types.Module{
					Variables: make([]*types.Variable, 0),
					Outputs:   make([]*types.Output, 0),
					Resources: make([]*types.Resource, 0),
				}
			}

			for _, block := range body.Blocks() {
				switch block.Type() {
				case "variable":
					modulesMap[moduleName].Variables = append(modulesMap[moduleName].Variables, p.parseVariable(block))
				case "output":
					modulesMap[moduleName].Outputs = append(modulesMap[moduleName].Outputs, p.parseOutput(block))
				case "resource":
					modulesMap[moduleName].Resources = append(modulesMap[moduleName].Resources, p.parseResource(block))
				}
			}

			return nil
		})
	if err != nil {
		p.logger.ErrorWithError("failed on scraping root folder please make sure path is correct", err)

		return nil, err
	}

	return modulesMap, nil
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
			typeTokens = append(typeTokens, v.Expr().BuildTokens(nil)...)
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
