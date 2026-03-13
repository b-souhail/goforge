package generate

import (
	"fmt"
	"goforge/internal/models"
	"os"
	"path/filepath"
)

func Scaffold(baseDir string, layers []models.Layer) error {
	for _, layer := range layers {
		layerPath := filepath.Join(baseDir, layer.Name)
		if err := os.MkdirAll(layerPath, 0755); err != nil {
			return fmt.Errorf("cannot create folder '%s': %w", layer.Name, err)
		}
		fmt.Printf("%s layer created\n", layer.Name)
		for _, dir := range layer.Dirs {
			subPath := filepath.Join(layerPath, dir)
			if err := os.MkdirAll(subPath, 0755); err != nil {
				return fmt.Errorf("cannot create '%s': %w", subPath, err)
			}
			fmt.Printf("\t%s directory created\n", dir)
		}
	}
	return nil
}
