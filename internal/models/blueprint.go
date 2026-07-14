package models

type Config struct {
	Path         string     `yaml:"path"`
	Name         string     `yaml:"name"`
	Modules      []string    `yaml:"modules"`
	Layers       []string    `yaml:"layers"`
	Resources    []Resource `yaml:"resources"`
}
