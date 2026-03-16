package models

type architectures struct {
	Layers []Layer
}
// 
var Architectures = map[string]architectures{
	"clean": {
		Layers: []Layer{
			{Name: "domain", Dirs: []string{"entity", "repository"}},
			{Name: "application", Dirs: []string{"dtos", "usecases"}},
			{Name: "infrastructure", Dirs: []string{"db", "repository"}},
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

// falls back to clean if the architecture is unknown. (available mvc clean)
func GetLayers(arch string) []Layer {
	if arch, exist := Architectures[arch]; exist {
		return arch.Layers
	}
	return Architectures["clean"].Layers
}
