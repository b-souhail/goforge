package cmd

import (
	"fmt"
	"goforge/internal/architecture"
	"goforge/internal/models"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var modules, arch string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new GoForge project",
	Long: `Initialize a new project using the GoForge architecture template.

This command creates the base project structure and optionally
generates initial modules.

The generated project contains the default architectural layers
used by GoForge such as domain,delivery, application and infrastructure.

Examples:
X = stil not implementef 

  goforge init my-app

  goforge init my-app --modules users,posts,comments X

  goforge init my-app -m users,posts,comments X

If no project name is provided, a default name will be used.

`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := "my-app" //default project name if no name provided
		projectArch := "clean"  //default arch if no arch selected

		if len(args) > 0 {
			projectName = args[0]
		}
		projectPath, _ := os.Getwd()
		projectPath = filepath.Join(projectPath, projectName)
		if err := os.MkdirAll(projectPath, 0755); err != nil {
			fmt.Println(fmt.Errorf("dir already exists : %v", err))  /// traitement derr nedded
			return
		}

		if err := os.Chdir(projectName); err != nil {   /// traitement derr  
			fmt.Println(err)  
			return
		}

		if _, err := os.Stat("goforge.yaml"); err == nil {
			fmt.Println(fmt.Errorf("config already exists : %v", err))  /// traitement derr needed
			return
		}
		config := models.Config{Path: projectPath, Architecture: projectArch, Name: projectName, Modules: modules}

		if err := architecture.CreateConfig(config); err != nil {
			fmt.Println(err)  /// traitement derr needed 
			return
		}

		fmt.Println("Project initialized  goforge.yaml ready for being configurated ")
		/// add inteligent cmd also set up conv neeeded conversastion setup
	},

}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP( // cmd ready for being used (need to reimplement )
		&modules,
		"modules",
		"m",
		"",
		"Modules to generate (comma separated: user,post,like)",
	)
	// initCmd.Flags().StringVarP(&arch, "arch", "a", "", "") //ignore 
	//add  func for seting up mvc instead of mvc   \\\\\\\\\\\\ez impl//////
}
