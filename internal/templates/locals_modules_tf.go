package templates

type LocalsModulesTF struct {
	Module map[string]*ModuleData
}

type MainModuleTF struct {
	Module   map[string]*MainModuleData
	RootPath string
}

type ComplexVariableData map[string]string

type ModuleData struct {
	SimpleLocals map[string]string
	MapLocals    map[string]ComplexVariableData
}

type MainModuleData struct {
	*ModuleData
	RequiredFields map[string]string
}
