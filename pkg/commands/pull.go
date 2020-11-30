package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/pull"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "pull",
	Short:                 "Pull all repos that don't have any non-pushed changes",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/#-fgh-pull",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			secrets     = configuration.GetSecrets()
			config      = configuration.GetConfig()
			clonedRepos = reposBasedOffCustomPath(cmd, config)
		)
		pull.PullRepos(secrets, clonedRepos)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	addCustomPathFlag(pullCmd)
}
