package generate

import (
	"bytes"
	"fmt"
	"goforge/internal/models"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v2"
)

func CreateYaml(cfg models.Config) error {
	yamlPath := filepath.Join(cfg.Path, "goforge.yaml")
	file, err := os.Create(yamlPath)
	if err != nil {
		return fmt.Errorf("create yaml file: %w", err)}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	layers, err := models.GetLayers(cfg.Architecture)
	if err != nil {
		return err
	}

	cfg.Layers = layers

	return encoder.Encode(cfg)
}

func CreateMain(cfg models.Config) error {
	cmdDir := filepath.Join(cfg.Name, "cmd")
	if err := os.MkdirAll(cmdDir, os.ModePerm); err != nil {
		return err
	}
	data, err := templatesFS.ReadFile("templates/cmd/main.go.tmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("main").Parse(string(data))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]string{
		"AppName": cfg.Name,
	})
	if err != nil {
		return err
	}
	mainPath := filepath.Join(cmdDir, "main.go")
	formatted, err := imports.Process(mainPath, buf.Bytes(), nil)
	if err != nil {
		return err
	}
	return os.WriteFile(mainPath, formatted, 0644)
}

func InitGoModules(cfg models.Config) error {
	cmd := exec.Command("go", "mod", "init", cfg.Name)
	cmd.Dir = cfg.Name
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go modules init failed: %v - %s", err, string(out))
	}
	return nil
}

func GenerateBase(config models.Config) error {

	if err := CreateYaml(config); err != nil { //ok
		return err
	}
	if err := CreateMain(config); err != nil { //not ok //remove magic number && set the errors
		return err
	}
	if err := InitGoModules(config); err != nil {//ok
		return err
	}

	return nil
}