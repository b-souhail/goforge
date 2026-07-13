package utils

import (
	"fmt"
	"goforge/internal/models"
	"os/exec"
)

//nope

func InitGoModules(cfg *models.Config) error {
	cmd := exec.Command("go", "mod", "init", cfg.Name)
	cmd.Dir = cfg.Name
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go modules init failed: %v - %s", err, string(out))
	}
	return nil
}
func TidyModules(cfg *models.Config) error {
	return nil
}
