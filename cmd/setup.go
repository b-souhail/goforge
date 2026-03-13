package cmd

import (
	"fmt"
	"goforge/internal/generate"
	"goforge/internal/models"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var setupCmd = &cobra.Command{
	Use:   "setup [config-path]",
	Short: "Scaffold project layers from goforge.yaml",
	Long: `Setup reads goforge.yaml and creates all layers with their sub-folders.
 
The yaml is the source of truth: if you add extra dirs or layers manually
in the file, setup will create them too. 
no flags needed.
 
Examples:
  goforge setup                          // reads ./goforge.yaml
  goforge setup ./myproject/goforge.yaml // reads a config at a given path

  `,
	Run: func(cmd *cobra.Command, args []string) {
		configPath := "goforge.yaml"
		if len(args) > 0 {
			configPath = args[0]
		}

		data, err := os.ReadFile(configPath)
		if err != nil {
			data, err = os.ReadFile(configPath + "/goforge.yaml")
			if err != nil {
				fmt.Printf("goforge.yaml not found\n")
				fmt.Println("Run: goforge init <project-name>\nOr: goforge setup ./path/goforge.yaml")
				return
			}
			configPath = configPath + "/goforge.yaml"
		}

		var config models.Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			fmt.Printf("Failed to parse '%s': %v\n", configPath, err)
			return
		}

		baseDir := filepath.Dir(configPath)
		if err := generate.Scaffold(baseDir, config.Layers); err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
