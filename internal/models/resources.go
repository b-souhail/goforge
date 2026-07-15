package models

type Resource struct {
	Name   string         `yaml:"name"`
	Params map[string]any `yaml:"params,omitempty"`
}

type GenerateFile struct {
	Template string
	Output   string
}

type ResourceDefinition interface {
	Name() string
	Questions() []*Question
	BuildConfig(Answers) Resource
	Files(Res Resource) []GenerateFile
	Data(cfg Config,Res Resource,) (any,any)
}
