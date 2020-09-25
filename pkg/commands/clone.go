package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone",
	Short:                 "Clone a repository",
	Run: func(cmd *cobra.Command, args []string) {
		secrets := configure.SecretsOutline{
			PAT:      "HERE IS THE PERSONAL ACCESS TOKEN",
			Username: "Matt-Gleich",
		}
		clone.GetRepository(secrets, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
