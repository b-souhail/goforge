package cmd

import (
	"fmt"
	"goforge/internal/generator"
	"goforge/internal/utils"
	"strings"

	"github.com/spf13/cobra"
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

  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := utils.ReadBlueprint("zz/" + utils.BlueprintFileName)
		if err != nil {
			return fmt.Errorf("error reading blueprint: %w", err)
		}

		if err := utils.InitGoModules(config.Name); err != nil {
			return fmt.Errorf("error Go modules: %w", err)
		}

		for _, module := range config.Modules {
			files := generator.ModuleFiles(module)
			for _, file := range files {
				if err := generator.Generate(file, *config, map[string]any{
					"Module":     strings.ToUpper(string(module[0])) + module[1:],
					"Receiver":   strings.ToLower(string(module[0])),
					"ModulePath": config.Name,
				}); err != nil {
					fmt.Printf("erreur génération %s: %v\n", file.Output, err)
				}
			}

		}

		// for _, v := range config.Resources {
		// 	definition := resources.Registry[v.Name]
		// 	files := definition.Files(v)
		// 	for _, file := range files {
		// 		generator.Generate(file, *config, map[string]any{})
		// 	}
		// }

		return utils.TidyModules(config)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
