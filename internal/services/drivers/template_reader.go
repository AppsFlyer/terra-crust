package drivers

type TemplateReader interface {
	GetRemoteModulesFromTemplate(templatePath string) (map[string]string, error)
}
