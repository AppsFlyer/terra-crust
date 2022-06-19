package types

type Resource struct {
	Type string `hcl:"type,label"`
	Name string `hcl:"name,label"`
}
