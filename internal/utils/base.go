package utils

import (
	"fmt"
	"goforge/internal/models"
	"os/exec"
)

//nope

func InitGoModules(name string) error {
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = name
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go modules init failed: %v - %s", err, string(out))
	}
	return nil
}
func TidyModules(cfg *models.Config) error {
	return nil
}
