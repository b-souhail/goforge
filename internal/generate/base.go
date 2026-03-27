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
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	cfg.Layers = models.GetLayers(cfg.Architecture)
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
	cmd := exec.Command("go", "mod", "init",cfg.Name)
	cmd.Dir = cfg.Name
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod init error: %v - %s", err, string(out))
	}
	return nil
}
