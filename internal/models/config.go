package models

type Config struct {
	Path         string   `yaml:"path"`
	Name         string   `yaml:"name"`
	Architecture string   `yaml:"architecture"`
	Modules      string   `yaml:"modules"`
	Layers       []string `yaml:"layers"`
	Plugins      string `yaml:"plugins"`
}
