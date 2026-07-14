package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goforge",
	Short: "CLI to scaffold and manage layered Go backend architectures",
	Long: `GoForge is a command-line tool designed to scaffold and evolve
layered backend architectures in Go.

It generates project structures, modules and architectural components
based on a predefined architecture pattern.

GoForge focuses on developer productivity by automating repetitive
tasks .

This tool is intended for developers who want to build clean,
structured backend services quickly while keeping full control
over their code.

Examples:
X : no usage until now implementation gonna be done later thx

  goforge init [projectName]   #by default the name is my-app 
  cd [projectName]
  goforge setup ./goforge.yaml #no need to specify the yaml file excepte if you re not in the dir

  goforge init my-app --modules users,posts X

Use "goforge --help" for more information about a command`,
}

func Execute() error {
	err := rootCmd.Execute()
	return err
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
