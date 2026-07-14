package utils

import (
	"slices"
	"fmt"
	"goforge/internal/models"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// dlta
const BlueprintFileName = "blueprint.yaml" //dlta
var Layers = []string{"domain","application","infrastructure","delivery",
}

func ReadBlueprint(FileName string) (*models.Config, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("File not found: %s", FileName)
	}
	var config models.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("Failed to parse '%s': %v\n", FileName, err)
	}

	return &config, nil
}

func SaveBlueprint(cfg *models.Config) error {
	path := filepath.Join(cfg.Path,cfg.Name, BlueprintFileName)

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating blueprint file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	defer encoder.Close()

	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("encode blueprint: %w", err)
	}

	return nil
}

func CreateBlueprint(cfg *models.Config) error {
	file, err := os.Create(filepath.Join(cfg.Path,cfg.Name ,BlueprintFileName))
	if err != nil {
		return fmt.Errorf("create yaml file: %w", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	cfg.Layers = Layers

	return encoder.Encode(cfg)
}

func HasResource(cfg *models.Config, name string) bool {
	for _, r := range cfg.Resources {
		if r.Name == name {
			return true
		}
	}
	return false
}
func HasModule(config *models.Config, name string) bool {
    return slices.Contains(config.Modules, name)
}