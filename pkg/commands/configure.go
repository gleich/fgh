package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure fgh with an interactive prompt.",
	Long:  "Before using fgh you need to configure it by running fgh configure. When it asks you for the GitHub PAT (personal access token) just go to [https://github.com/settings/tokens/new](https://github.com/settings/tokens/new) and create a new token with the repo box check off.",
	Run: func(cmd *cobra.Command, args []string) {
		regularConfig := configure.AskQuestions()
		secretConfig := configure.AskSecretQuestions()
		configure.WriteConfiguration(secretConfig, regularConfig)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
