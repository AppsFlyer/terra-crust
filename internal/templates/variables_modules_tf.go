package templates

type VariblesModuleList map[string]*VariablesModulesTF

type VariablesModulesTF struct {
	ModuleName        string
	Description       string
	ObjectTypeMapping map[string]string
	DefaultValues     map[string]string
}
