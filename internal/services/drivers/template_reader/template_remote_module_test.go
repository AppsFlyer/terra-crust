package template_reader_test

import (
	version_control "github.com/AppsFlyer/terra-crust/internal/services/drivers/version_control"
	"github.com/go-test/deep"
	"testing"

	logger "github.com/AppsFlyer/go-logger"
	tmplReader "github.com/AppsFlyer/terra-crust/internal/services/drivers/template_reader"
)

var result = map[string]*version_control.RemoteModule{
	"terra-crust":    {Name: "terra-crust", Url: "https://github.com/AppsFlyer/terra-crust", Version: "", Path: ""},
	"naming":         {Name: "naming", Url: "https://github.domain.com/test/terraform/modules/naming.git", Version: "0.2.1", Path: "modules/naming"},
	"otel-collector": {Name: "otel-collector", Url: "https://github.com/streamnative/terraform-helm-charts.git", Version: "v0.2.1", Path: "modules/otel-collector"},
	"iam-account":    {Name: "iam-account", Path: "modules/iam-account", Url: "https://github.com/terraform-aws-modules/terraform-aws-iam.git", Version: ""},
	"zones":          {Name: "zones", Path: "modules/zones", Url: "https://github.com/terraform-aws-modules/terraform-aws-route53.git", Version: ""},
}

func TestGetRemoteModulesFromTemplate(t *testing.T) {
	log := logger.NewSimple()
	templateReader := tmplReader.InitTemplateRemoteModule(log)

	modules, err := templateReader.GetRemoteModulesFromTemplate("./test.tmpl")
	if err != nil {
		t.Errorf("failed to extract sources from template %s", err.Error())
	}

	if len(modules) != 5 {
		t.Errorf("Expected 5 modules found : %d", len(modules))
	}

	if diff := deep.Equal(modules, result); diff != nil {
		t.Errorf("expected result to be equal, result is different : %s", diff)
	}
}

func TestBadPath(t *testing.T) {
	log := logger.NewSimple()
	templateReader := tmplReader.InitTemplateRemoteModule(log)

	_, err := templateReader.GetRemoteModulesFromTemplate("./main.tmpl")
	if err == nil {
		t.Errorf("expected to have error for no file found")
	}

}
