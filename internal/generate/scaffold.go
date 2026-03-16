package generate

import (
	"embed"
	"fmt"
	"goforge/internal/models"
)

//go:embed templates/clean
var templatesFS embed.FS

func ScaffoldConfig(cfg models.Config) error {
    fmt.Println(cfg.Name)

	return nil
}