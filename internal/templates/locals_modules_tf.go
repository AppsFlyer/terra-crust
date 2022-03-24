package templates

type LocalsModulesTF struct {
	Module map[string]*ModuleData
}

type ComplexVariableData map[string]string

type ModuleData struct {
	SimpleLocals map[string]string
	MapLocals    map[string]ComplexVariableData
}
