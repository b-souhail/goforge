package generator

import (
	"bytes"
	"embed"
	"fmt"
	"goforge/internal/models"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/tools/imports"
)

//go:embed templates/*
var templatesFS embed.FS

// func Generate(file models.GenerateFile, config models.Config, params any) error {

// 	data, err := templatesFS.ReadFile("templates/" + file.Template)
// 	if err != nil {
// 		return err
// 	}

// 	tmpl, err := template.New(filepath.Base(file.Template)).Parse(string(data))
// 	if err != nil {
// 		return err
// 	}

// 	var buf bytes.Buffer

// 	if err := tmpl.Execute(&buf, params); err != nil {
// 		return err
// 	}

// 	output := filepath.Join(config.Path, config.Name, file.Output)

// 	if err := os.MkdirAll(filepath.Dir(output), 0755); err != nil {
// 		return err
// 	}

// 	formatted, err := imports.Process(output, buf.Bytes(), &imports.Options{Comments: true})
// 	if err != nil {
// 		return err
// 	}

// 	if err := os.WriteFile(output, formatted, 0644); err != nil {
// 		return err
// 	}

// 	return nil
// }

func Generate(file models.GenerateFile, config models.Config, params any) error {

	data, err := templatesFS.ReadFile("templates/" + file.Template)
	if err != nil {
		return err
	}

	tmpl, err := template.New(filepath.Base(file.Template)).Parse(string(data))
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

	formatted, fmtErr := imports.Process(output, buf.Bytes(), &imports.Options{})
	if fmtErr != nil {
		// goimports a échoué (souvent une erreur de syntaxe dans le template
		// exécuté) : on écrit quand même le fichier brut pour que le
		// contenu généré reste visible et débuggable, au lieu de disparaître.
		if err := os.WriteFile(output, buf.Bytes(), 0644); err != nil {
			return err
		}
		return fmt.Errorf("%s généré mais non formaté (invalide): %w", output, fmtErr)
	}

	if err := os.WriteFile(output, formatted, 0644); err != nil {
		return err
	}

	return nil
}




// //go:embed templates/*
// var templatesFS embed.FS

// func ScaffoldConfig(cfg *models.Config) error {
// 	modules := strings.Split(cfg.Modules[0], ",") //ToDo

// 	for _, layer := range cfg.Layers {
// 		lPath := filepath.Join(cfg.Name, "internal", layer.Name)
// 		if err := os.MkdirAll(lPath, 0755); err != nil {
// 			return err
// 		}
// 		templateDir := filepath.Join("templates", "clean", layer.Name)

// 		entries, err := templatesFS.ReadDir(templateDir)
// 		if err != nil {
// 			fmt.Println("no templates  for " + layer.Name)
// 			continue
// 		}

// 		for _, entry := range entries {
// 			subDir := strings.TrimSuffix(entry.Name(), ".go.tmpl")
// 			subPath := filepath.Join(lPath, subDir) // destination where to generate the tmpl
// 			for _, module := range modules {
// 				if module == "" {
// 					continue
// 				}

// 				if err := generateFile(cfg, templateDir, subPath, entry.Name(), module); err != nil {
// 					return err
// 				}

// 			}

// 		}

// 	}
// 	return nil
// }

// func generateFile(cfg *models.Config, templateDir, destDir, tmplName, module string) error {
// 	tmplPath := filepath.Join(templateDir, tmplName)

// 	data, err := templatesFS.ReadFile(tmplPath)
// 	if err != nil {
// 		return err
// 	}
// 	tmpl, err := template.New(tmplName).Parse(string(data))
// 	if err != nil {
// 		return err
// 	}

// 	tmplN := strings.Split(tmplName, ".")

// 	fileName := fmt.Sprintf("%s_%s.go", module, tmplN[0])
// 	fullPath := filepath.Join(destDir, fileName)

// 	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
// 		return err
// 	}

// 	fLetter := strings.ToUpper(string(module[0]))
// 	format := fLetter + module[1:]

// 	var buf bytes.Buffer
// 	err = tmpl.Execute(&buf, map[string]string{
// 		"Module":     format,
// 		"Receiver":   strings.ToLower(string(module[0])),
// 		"ModulePath": cfg.Name,
// 	})
// 	if err != nil {  
// 		return err
// 	}
// 	if err := os.WriteFile(fullPath, buf.Bytes(), 0644); err != nil {
// 		return err
// 	}
// 	opts := &imports.Options{}
// 	formatted, err := imports.Process(fullPath, buf.Bytes(), opts)
// 	if err != nil {
// 		return err
// 	}
// 	return os.WriteFile(fullPath, formatted, 0644)
// }
