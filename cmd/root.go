package cmd

import (
	"os"

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

  goforge init my-app X
  goforge init my-app --modules users,posts X

  goforge add module comments X
  goforge add usecase create-user X

Use "goforge --help" for more information about a command`,



	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
