package models

import (
	"os"

	"gopkg.in/yaml.v2"
)

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
