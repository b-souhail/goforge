package resources

import (
	"fmt"
	"goforge/internal/models"
)

type HTTP struct{}

func (HTTP) Name() string {
	return "http"
}

func (HTTP) Questions() []*models.Question {
	return []*models.Question{
		{
			Key:  "middleware",
			Text: "Do you want customized middlewares?",

			Next: &models.Question{
				Key:     "middlewares",
				Text:    "Select middleware",
				Options: []string{"cors", "limiter","secure_header"},
			},
		},
		{
			Key:  "websocket",
			Text: "Do you want websocket?",

			Next: &models.Question{
				Key:  "hub",
				Text: "Enable websocket hub?",
			},
		},
	}
}

func (HTTP) Files(res models.Resource) []models.GenerateFile {
	files := []models.GenerateFile{
		{
			Template: "resources/http/routes.go.tmpl",
			Output:   "internal/delivery/http/routes.go",
		},{
			Template: "resources/http/server.go.tmpl",
			Output:   "internal/delivery/http/server.go",
		},
	}

	if middlewares, ok := res.Params["middlewares"].([]any); ok {
		for _, mw := range middlewares {
			name := mw.(string)

			files = append(files, models.GenerateFile{
				Template: fmt.Sprintf("resources/http/middlewares/%s.go.tmpl", name),
				Output:   fmt.Sprintf("internal/delivery/http/middlewares/%s.go", name),
			})
		}
	}

	if res.Params["websocket"].(bool) {
		temp := models.GenerateFile{
			Template: "resources/http/websockets/ws.go.tmpl",
			Output:   "internal/delivery/websockets/ws.go",
		}
		files = append(files, temp)
	}
	if res.Params["hub"] != nil && res.Params["hub"].(bool) {
		temp := []models.GenerateFile{
			{
				Template: "resources/http/websockets/hub.go.tmpl",
				Output:   "internal/delivery/websockets/hub.go",
			},
			{
				Template: "resources/http/websockets/connection.go.tmpl",
				Output:   "internal/delivery/websockets/connection.go",
			},
		}
		files = append(files, temp...)
	}

	return files
}

func (HTTP) BuildConfig(a models.Answers) models.Resource {
	params := make(map[string]any)

	if a["websocket"].(bool) {
		params["websocket"] = a["websocket"].(bool)
		params["hub"] = a["hub"].(bool)
	} else {
		params["websocket"] = a["websocket"].(bool)
	}
	if a["middleware"].(bool) {
		params["middleware"] = a["middleware"].(bool)
		params["middlewares"] = a["middlewares"].([]string)
	} else {
		params["middleware"] = a["middleware"].(bool)
	}

	return models.Resource{
		Name:   "http",
		Params: params,
	}
}

type HTTPData struct {
	Modules    []string
	Websocket  bool
	Hub        bool
	Middleware middleware
}

type middleware struct {
	ReturnExp string
}

func (HTTP) Data(config models.Config, res models.Resource) (any, any) {

	data := HTTPData{
		Modules:   config.Modules,
		Websocket: res.Params["websocket"].(bool),
		Hub:       res.Params["hub"].(bool),
	}

	expr := "mux"

	if res.Params["middleware"].(bool) {
		for _, mw := range res.Params["middlewares"].([]string) {
			switch mw {
			case "secure_headers":
				expr = fmt.Sprintf("MW.SecureHeaders(%s)", expr)

			case "cors":
				expr = fmt.Sprintf("MW.CORSMiddleware(%s)", expr)
			}
		}
	}
	data.Middleware.ReturnExp = expr

	return data, "HTTP"
}
