package cmd

import (
	"fmt"
	"goforge/internal/generate"
	"goforge/internal/models"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var archFlag string
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new GoForge project",
	Long: `Initialize a new project using the GoForge architecture template.

This command creates the base project structure and optionally
generates initial modules.

The generated project contains the default architectural layers
used by GoForge such as domain,delivery, application and infrastructure or 
MVC architecture based on Model, View and Controller layers.

Examples:

  goforge init                        # creates my-project/ folder && goforge.yaml file whit clean architecture
  goforge init myproject --arch mvc   # creates myproject/ folder && goforge.yaml file  with mvc architecture

 }

`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		//fallback to projectName if no name provided
		projectName := "my-project" //default name

		if len(args) == 1 {
			projectName = args[0]
		}

		projectPath, _ := os.Getwd()
		projectPath = filepath.Join(projectPath, projectName)
		
		if _, err := os.Stat(projectPath); err == nil {
			return fmt.Errorf("directory %s already exists\n", projectName)
		}

		const dirPerm = 0o755
		if err := os.MkdirAll(projectPath, dirPerm); err != nil {
			return fmt.Errorf("create directory: %w \n", err)
		}

		config := models.Config{Path: projectPath, Architecture: archFlag, Name: projectName}

		if err := generate.GenerateBase(config); err != nil {
			return fmt.Errorf("generate base: %w", err)
		}

		cmd.Printf(`Project %s created successfully.

Next steps:
cd %s
goforge setup

You can also edit goforge.yaml file before running setup command.
Use "goforge setup --help" to see all available commands.
`, projectName, projectName)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&archFlag, "arch", "a", "clean", "Architecture type: clean, mvc")
}
