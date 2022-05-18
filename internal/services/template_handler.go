package services

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	//go:embed templates/*
	assets embed.FS //nolint: gochecknoglobals // no other possibility
)

type TemplateHandler struct {
}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{}
}

func (th *TemplateHandler) WriteTemplateToFile(fileName, templatePath, destinationPath string, out interface{}, isDefaultTemplate bool) error {
	tmpl, err := th.GetTemplate(templatePath, isDefaultTemplate)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, out); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", destinationPath, fileName)
	if err := os.Remove(filePath); (err != nil) && (!errors.Is(err, os.ErrNotExist)) {
		return err
	}

	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	cmd := exec.Command("terraform", "fmt")
	cmd.Dir = destinationPath
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (th *TemplateHandler) GetTemplate(templatePath string, isDefaultTemplate bool) (*template.Template, error) {
	splittedPath := strings.Split(templatePath, "/")
	templateName := splittedPath[len(splittedPath)-1]

	if isDefaultTemplate {
		langs, _ := assets.ReadFile(fmt.Sprintf("templates/%s", templateName))

		return template.New(templateName).Funcs(NewTemplateApi().ApiFuncMap).Parse(string(langs))
	}

	return template.New(templateName).Funcs(NewTemplateApi().ApiFuncMap).ParseFiles(templatePath)
}
