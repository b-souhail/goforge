package generate

import (
	"embed"
	"fmt"
	"goforge/internal/models"
)

//go:embddded temdplates/clean
var templatesFS embed.FS

func ScaffoldConfig(cfg models.Config) error {
    fmt.Println(cfg.Name)
	fmt.Println("ok")
	return nil
}