/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goforge/internal/resources"
	"goforge/internal/utils"
	"strings"

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

		config, err := utils.ReadBlueprint("zz/" + utils.BlueprintFileName)
		if err != nil {
			return fmt.Errorf("error reading blueprint: %w", err)
		}
		if modulesFlag != "" {
			if !utils.HasModule(config, modulesFlag) {
				modulesFlag = strings.ToLower(modulesFlag)
				config.Modules = append(config.Modules, modulesFlag)
			}
			return utils.SaveBlueprint(config)

		}
		resource, ok := resources.Registry[args[0]]
		if !ok {
			return fmt.Errorf("error unknown resource")
		}
		if utils.HasResource(config, resource.Name()) {
			return fmt.Errorf("resource %q already exists", resource.Name())
		}

		answers, err := utils.AskQuestions(resource.Questions())
		if err != nil {
			return fmt.Errorf("error asking questions: %w", err)
		}

		fmt.Println("befor apend",config.Resources)
		config.Resources = append(config.Resources, resource.BuildConfig(answers))
		fmt.Println("after apend",config.Resources)

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


// func SaveBlueprint(cfg *models.Config) error {
// 	path := filepath.Join(cfg.Path, BlueprintFileName)

// 	file, err := os.Create(path)
// 	if err != nil {
// 		return fmt.Errorf("error creating blueprint file: %w", err)
// 	}
// 	defer file.Close()

// 	encoder := yaml.NewEncoder(file)
// 	encoder.SetIndent(2)
// 	defer encoder.Close()

// 	if err := encoder.Encode(cfg); err != nil {
// 		return fmt.Errorf("encode blueprint: %w", err)
// 	}

// 	return nil
// }