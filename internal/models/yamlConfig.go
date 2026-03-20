package models

type Layer struct {
	Name string   `yaml:"name"`
	Dirs []string `yaml:"dirs,omitempty"` 
	//todo : setup subdirectory 
}

type Config struct {
	Path         string  `yaml:"-"` // runtime
	Name         string  `yaml:"name"`
	Architecture string  `yaml:"architecture"`
	Modules      string  `yaml:"modules"`
	Layers       []Layer `yaml:"layers"`
	Plugins      string  `yaml:"plugins,omitempty"`  //maybe for V1.0.0
}
