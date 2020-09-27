package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone <OWNER/REPO>",
	Short:                 "Clone a repository",
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		secrets := configuration.GetSecrets()
		clone.GetRepository(secrets, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
