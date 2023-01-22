package template_reader_test

import (
	"github.com/go-test/deep"
	"testing"

	logger "github.com/AppsFlyer/go-logger"
	tmplReader "github.com/AppsFlyer/terra-crust/internal/services/drivers/template-reader"
)

var result = map[string]string{
	"moduleName":        "https://gitlab.domain.com/sub/sub/modules/moduleName",
	"moduleWithVersion": "https://github.com/AppsFlyer/terra-crust/random/long/path//moduleWithVersion",
	"terra-crust":       "https://github.com/AppsFlyer/terra-crust",
	"naming":            "https://github.domain.com/test/terraform/modules/naming.git",
}

func TestGetRemoteModulesFromTemplate(t *testing.T) {
	t.Parallel()
	log := logger.NewSimple()
	templateReader := tmplReader.InitTemplateRemoteModule(log)

	modules, err := templateReader.GetRemoteModulesFromTemplate("./main.tmpl")
	if err != nil {
		t.Errorf("failed to extract sources from template %s", err.Error())
	}

	if len(modules) != 4 {
		t.Errorf("Expected 4 modules found : %d", len(modules))
	}

	if diff := deep.Equal(modules, result); diff != nil {
		t.Errorf("expected result to be equal, result is different : %s", diff)
	}

}
