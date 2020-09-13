package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure fgh",
	Run: func(cmd *cobra.Command, args []string) {
		configure.AskQuestions()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
