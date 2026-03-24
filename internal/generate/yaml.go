package generate

import (
	"goforge/internal/models"
	"os"
	"path/filepath"

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
