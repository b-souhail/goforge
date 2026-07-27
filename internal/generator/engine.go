package generator

import (
	"bytes"
	"embed"
	"fmt"
	"goforge/internal/models"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
)

//go:embed templates/*
var templatesFS embed.FS

var Funcs = template.FuncMap{
	"Capitalize": func(s string) string {
		if s == "" {
			return s
		}
		return strings.ToUpper(s[:1]) + s[1:]
	},
	"Lower": func(s string) string {
		if s == "" {
			return s
		}
		return strings.ToLower(s)
	},
	"Receiver": func(s string) string {
		if s == "" {
			return s
		}
		return strings.ToLower(string(s[0]))
	},
}

func Generate(file models.GenerateFile, config models.Config, params any) error {

	data, err := templatesFS.ReadFile("templates/" + file.Template)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filepath.Base(file.Template)).Funcs(Funcs).Parse(string(data))
	if err != nil {
		return err
	}
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, params); err != nil {
		return err
	}
	output := filepath.Join(config.Path, config.Name, file.Output)

	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
		return err
	}

	formatted, fmtErr := imports.Process(output, buf.Bytes(), &imports.Options{
		Comments: true,
	})
	if fmtErr != nil {
		if err := os.WriteFile(output, buf.Bytes(), 0644); err != nil {
			fmt.Println(err)
			return err
		}
		return fmt.Errorf("%s généré mais non formaté (invalide): %w", output, fmtErr)
	}

	if err := os.WriteFile(output, formatted, 0644); err != nil {
		return err
	}

	return nil
}
