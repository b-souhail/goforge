package models

import "fmt"

type Architectures struct {
	Layers []Layer
}

var architectures = map[string]Architectures{
	"clean": {
		Layers: []Layer{
			{Name: "domain", Dirs: []string{"entity", "repository"}},
			{Name: "application", Dirs: []string{"dtos", "usecases"}},
			{Name: "infrastructure", Dirs: []string{"repository"}},
			{Name: "delivery", Dirs: []string{}},
		},
	},
	"mvc": {
		Layers: []Layer{
			{Name: "models", Dirs: []string{}},
			{Name: "views", Dirs: []string{}},
			{Name: "controllers", Dirs: []string{}},
		},
	},
}

func GetLayers(name string) ([]Layer, error) {
	architecture, ok := architectures[name]
	if !ok {
		return nil, fmt.Errorf("architecture %s not available or unknown", name)
	}
	return architecture.Layers, nil
}
