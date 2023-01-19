package drivers

type TemplateReader interface {
	GetRemoteModulesFromCache() map[string]string
	GetRemoteModulesFromTemplate(templatePath string) (map[string]string, error)
}
