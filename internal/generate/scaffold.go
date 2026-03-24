package generate

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

func ScaffoldConfig(c models.Config) error {
	modules := strings.Split(c.Modules, ",")

	for _, l := range c.Layers {
		lPath := filepath.Join(c.Name, l.Name)
		if err := os.MkdirAll(lPath, 0755); err != nil {
			return err
		}
		for _, dir := range l.Dirs {
			subDirPath := filepath.Join(lPath, dir)
			if err := os.MkdirAll(subDirPath, 0755); err != nil {
				return err
			}
		}
		templateDir := filepath.Join("templates", c.Architecture, l.Name)
		entries, err := templatesFS.ReadDir(templateDir)
		if err != nil {
			fmt.Println("no templates  for " + l.Name)
			continue
		}

		for _, entry := range entries {
			subDir := strings.TrimSuffix(entry.Name(), ".go.tmpl")
			subPath := filepath.Join(lPath, subDir) // destination where to generate the tmpl
			for _, module := range modules {
				if module == "" {
					continue
				}

				if err := generateFile(templateDir, subPath, entry.Name(), module); err != nil {
					return err
				}

			}

		}

	}
	return nil
}

func generateFile(templateDir, destDir, tmplName, module string) error {
	tmplPath := filepath.Join(templateDir, tmplName)
	data, err := templatesFS.ReadFile(tmplPath)
	if err != nil {
		
		return err
	}
	tmpl, err := template.New(tmplName).Parse(string(data))
	if err != nil {
				fmt.Println("okokok")

		return err
	}
	tmplN := strings.Split(tmplName, ".")

	fileName := fmt.Sprintf("%s_%s.go", module, tmplN[0])

	f, err := os.Create(filepath.Join(destDir, fileName))
	if err != nil {
				fmt.Println("okokok")

		return err
	}
	defer f.Close()

	fLetter := strings.ToTitle(string(module[0]))
	format := fmt.Sprintf("%s%s", fLetter, string(module[1:]))

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]string{
		"Module":   format,
		"Receiver": strings.ToLower(string(module[0])),
	})
	if err != nil {
		return err
	}
	formatted, err := imports.Process(filepath.Join(destDir, fileName), buf.Bytes(), nil)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(destDir, fileName), formatted, 0644)
}
