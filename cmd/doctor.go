package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Analyze project structure and detect architecture violations",
	Long: `The doctor command scans the project files and analyzes imports
to detect violations of the architecture rules.

It is mainly used to ensure that dependencies between layers respect
the expected boundaries.

Usage:

  goforge doctor

The command will recursively scan the project files and report
any invalid import relationship between layers.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("doctor give pills")
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}