package drivers

type VersionControl interface {
	CloneModules(modules map[string]string, destination string) error
}
