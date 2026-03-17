package models

type Layer struct {
	Name string   `yaml:"name"`
	Dirs []string `yaml:"dirs,omitempty"`
}

type Config struct {
	Path         string  `yaml:"-"`
	Name         string  `yaml:"name"`
	Architecture string  `yaml:"architecture"`
	Modules      string  `yaml:"modules"`
	Layers       []Layer `yaml:"layers"`
	Plugins      string  `yaml:"plugins,omitempty"`
}
