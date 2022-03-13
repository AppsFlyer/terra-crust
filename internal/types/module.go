package types

import (
	"fmt"
	"strings"
)

type Module struct {
	Name      string      `hcl:"name,label"`
	Variables []*Variable `hcl:"variable,block"`
	Outputs   []*Output   `hcl:"output,block"`
	Resources []*Resource `hcl:"resource,block"`
	Source    string      `hcl:"source,attr"`
}

func (m *Module) String() string {
	var s string
	s = fmt.Sprintf("Name: %s\nSource: %s\nVariables: \n", m.Name, m.Source)
	for _, v := range m.Variables {
		s += fmt.Sprintf("  %s", v.Name)
		s += fmt.Sprintf(" %s", strings.ReplaceAll(string(v.Type.Bytes()), " ", ""))
		s += fmt.Sprintf(" %s", v.Description)
		s += fmt.Sprintf(" %s\n", strings.ReplaceAll(string(v.Default.Bytes()), " ", ""))
	}
	s += "Outputs: \n"
	for _, v := range m.Outputs {
		s += fmt.Sprintf("  %s", v.Name)
		s += fmt.Sprintf(" %s", strings.ReplaceAll(string(v.Value.Bytes()), " ", ""))
		s += fmt.Sprintf(" %s\n", v.Description)
	}

	return s
}
