package cmd

import (
	"fmt"
	"goforge/internal/resources"
	"goforge/internal/utils"

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
		config, err := utils.ReadBlueprint("z/" + utils.BlueprintFileName)
		if err != nil {
			return fmt.Errorf("error reading blueprint: %w", err)
		}

		for _, v := range config.Resources {
			definition := resources.Registry[v.Name]
			res := definition.Files(v)
			for _, v := range res {
				fmt.Println(v.Output)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
