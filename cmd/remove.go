/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goforge/internal/resources"
	"goforge/internal/utils"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	RunE: func(cmd *cobra.Command, args []string) error {
	resource, ok := resources.Registry[args[0]]
		if !ok {
			return fmt.Errorf("error unknown resource")
		}

		cfg, err := utils.ReadBlueprint("z/" + utils.BlueprintFileName)
		if err != nil {
			return fmt.Errorf("error reading blueprint: %w", err)
		}

		if utils.HasResource(cfg, resource.Name()) {
			return fmt.Errorf("resource %q  exist", resource.Name())
		}
		// if utils.HasModule(cfg) {
		// 	return fmt.Errorf("resource %q  exist", resource.Name())
		// }
		
		fmt.Println("do not exist")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
