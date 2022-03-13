package types

import "github.com/hashicorp/hcl/v2/hclwrite"

type Output struct {
	Name        string          `hcl:"name,label"`
	Value       hclwrite.Tokens `hcl:"value,attr"`
	Description string          `hcl:"description,attr"`
}
