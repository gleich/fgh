package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure fgh with an interactive prompt",
	Run: func(cmd *cobra.Command, args []string) {
		config := configure.AskQuestions()
		configure.WriteConfiguration(config)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
