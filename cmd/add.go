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

var modulesFlag string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		resource, ok := resources.Registry[args[0]]
		if !ok {
			return fmt.Errorf("error unknown resource")
		}

		config, err := utils.ReadBlueprint("z/" + utils.BlueprintFileName)
		if err != nil {
			return fmt.Errorf("error reading blueprint: %w", err)
		}

		if utils.HasResource(config, resource.Name()) {
			return fmt.Errorf("resource %q already exists", resource.Name())
		}

		answers, err := utils.AskQuestions(resource.Questions())
		if err != nil {
			return fmt.Errorf("error asking questions: %w", err)
		}

		config.Resources = append(config.Resources, resource.BuildConfig(answers))

		return utils.SaveBlueprint(config)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(
		&modulesFlag,
		"modules",
		"m",
		"",
		"Modules to generate (comma separated: user,post,like)",
	)
}
