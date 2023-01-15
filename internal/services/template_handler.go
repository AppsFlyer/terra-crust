// Copyright 2022 AppsFlyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package services

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	logger "github.com/AppsFlyer/go-logger"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

var (
	//go:embed templates/*
	assets embed.FS //nolint: gochecknoglobals // no other possibility
)

type TemplateHandler struct {
	logger logger.Logger
}

func NewTemplateHandler(log logger.Logger) *TemplateHandler {
	return &TemplateHandler{
		logger: log,
	}
}

func (th *TemplateHandler) runTerraformFmt(path string) error {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		th.logger.Error("Failed installing terraform", err.Error())

		return err
	}

	workingDir := path
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		th.logger.Error("failed running NewTerraform", err.Error())

		return err
	}

	if err := tf.FormatWrite(context.Background()); err != nil {
		th.logger.Error("failed running Show", err.Error())

		return err
	}

	return nil
}

func (th *TemplateHandler) WriteTemplateToFile(fileName, templatePath, destinationPath string, out interface{}, isDefaultTemplate bool) error {
	tmpl, err := th.GetTemplate(templatePath, isDefaultTemplate)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, out); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", destinationPath, fileName)
	if err = os.Remove(filePath); (err != nil) && (!errors.Is(err, os.ErrNotExist)) {
		return err
	}

	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	if err := th.runTerraformFmt(destinationPath); err != nil {
		th.logger.Errorf("Failed running terraform FMT please make sure directory is correct, or external template if provided is correct")
	}

	return nil
}

func (th *TemplateHandler) GetTemplate(templatePath string, isDefaultTemplate bool) (*template.Template, error) {
	splittedPath := strings.Split(templatePath, "/")
	templateName := splittedPath[len(splittedPath)-1]

	apiFunc := NewTemplateAPI()
	if isDefaultTemplate {
		langs, _ := assets.ReadFile(fmt.Sprintf("templates/%s", templateName))

		return template.New(templateName).Funcs(*apiFunc.APIFuncMap).Parse(string(langs))
	}

	return template.New(templateName).Funcs(*apiFunc.APIFuncMap).ParseFiles(templatePath)
}
