package models

type Layer struct {
	Name string   `yaml:"name"`
	Dirs []string `yaml:"directory"`
}

var Architecture = []Layer{
	{Name: "domain"},
	{Name: "application"},
	{Name: "infrastructure"},
	{Name: "delivery"},
}

