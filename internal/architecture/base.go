package architecture

import (
	"goforge/internal/models"
	"os"

	"gopkg.in/yaml.v2"
)

func CreateConfig(config models.Config) error {
	layer := []string{"domain", "application", "delivery", "infrastructure"} // to impl need to read from directry adapte and setup
	config.Layers = layer
	config.Plugins = "" //ignore 

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile("goforge.yaml", data, 0644)
}
