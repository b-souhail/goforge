package http

import (
	"fmt"
	"goforge/internal/models"
)

type HTTPData struct {
	Modules    []string
	Websocket  bool
	Hub        bool
	Middleware middleware
}

type middleware struct {
	HandleFuncExp string
	HandlerExp    string
}

func (HTTP) Data(config models.Config, res models.Resource) any {

	data := &HTTPData{
		Modules:   config.Modules,
		Websocket: res.Params["websocket"].(bool),
		Hub:       res.Params["hub"].(bool),
	}

	if res.Params["middleware"].(bool) {
		data.Middleware.handlerExpression(data, toStrings(res.Params["middlewares"]))
		data.Middleware.handleFuncExpression(data, toStrings(res.Params["middlewares"]))
	}

	return data
}

func (m middleware) handlerExpression(data *HTTPData, middlewares []string) {
	var expr string = "mux"
	for _, mw := range middlewares {
		switch mw {
		case "secure_header":
			expr = fmt.Sprintf("mw.SecureHeader(%s)", expr)

		case "cors":
			expr = fmt.Sprintf("mw.CORS(%s)", expr)
		}
	}
	data.Middleware.HandlerExp = expr
}

func (m middleware) handleFuncExpression(data *HTTPData, middlewares []string) {

	var expr string = "%s"

	for _, mw := range middlewares {
		switch mw {
		case "limiter":
			expr = fmt.Sprintf("mw.Limiter(%s)", expr)
		}
	}

	data.Middleware.HandleFuncExp = expr

}

func toStrings(v any) []string {
	items, ok := v.([]any)
	if !ok {
		return nil
	}

	result := make([]string, 0, len(items))
	for _, item := range items {
		if s, ok := item.(string); ok {
			result = append(result, s)
		}
	}

	return result
}
