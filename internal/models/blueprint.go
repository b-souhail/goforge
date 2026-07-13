package models

type Config struct {
	Path         string     `yaml:"path"`
	Name         string     `yaml:"name"`
	Modules      string     `yaml:"modules"`
	Layers       []Layer    `yaml:"layers"`
	Resources    []Resource `yaml:"resources"`
}