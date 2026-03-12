package generate

import (
	"goforge/internal/models"
	"os"

	"gopkg.in/yaml.v2"
)

func CreateYaml(config models.Config) error {
	config.Layers = models.GetLayers(config.Architecture)
	config.Plugins = "" //ignore
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile("goforge.yaml", data, 0644)
}
