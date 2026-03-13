package cmd

import (
	"fmt"
	"goforge/internal/models"
	"goforge/internal/generate"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var modulesFlag, archFlag string
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new GoForge project",
	Long: `Initialize a new project using the GoForge architecture template.

This command creates the base project structure and optionally
generates initial modules.

The generated project contains the default architectural layers
used by GoForge such as domain,delivery, application and infrastructure.

Examples:

  goforge init                        # creates my-app/ goforge.yaml with clean arch
  goforge init myproject              # creates myproject/ goforge.yaml with clean arch
  goforge init myproject --arch mvc   # creates myproject/ goforge.yaml with mvc arch
  goforge init myproject -a mvc       # creates myproject/ goforge.yaml with mvc arch



  {//this cmd will be combined white setup cmd
  X = stil not implementef 

  goforge init my-app --modules users,posts,comments X
  goforge init my-app -m users,posts,comments X
  
 }

`,
	Run: func(cmd *cobra.Command, args []string) {

		projectName := "my-app" //default project name if no name provided
		if len(args) > 0 {
			projectName = args[0]
		}

		projectPath, _ := os.Getwd()
		projectPath = filepath.Join(projectPath, projectName)
		if err := os.MkdirAll(projectPath, 0755); err != nil {
			fmt.Println(fmt.Errorf("dir already exists : %v", err)) //todo\\
			return
		}

		if err := os.Chdir(projectName); err != nil {
			fmt.Println(err) //todo\\
			return
		}
		config := models.Config{Path: projectPath, Architecture: archFlag, Name: projectName, Modules: modulesFlag}

		if err := generate.CreateYaml(config); err != nil {
			fmt.Println(err) //todo\\
			return
		}

		fmt.Printf("Project %s ready for beign setup with %s architecture\n", projectName, archFlag)
		fmt.Printf("Next steps:\ngoforge setup %v\n", projectName)

		if modulesFlag != "" {
			    configFile := filepath.Join(projectPath, "goforge.yaml")

            config, err := config.ReadYaml(configFile)
            if err != nil {
                fmt.Println(err) ///\\\
            }
            if err := generate.Scaffold(projectPath, config.Layers); err != nil {
                fmt.Println(err) ///\\\\
                return
            }
        }
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&archFlag, "arch", "clean", "Architecture type: clean, mvc")
	initCmd.Flags().StringVarP(
		&modulesFlag,
		"modules",
		"m",
		"",
		"Modules to generate (comma separated: user,post,like)",
	)
}
