package models

import (
	"os"

	"gopkg.in/yaml.v2"
)

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

func (c *Config) ReadYaml(configFile string) (Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

//   - name: str
//     dirs:
//       -str1
//       -str2
//       -str2/str    //todo : implement\
