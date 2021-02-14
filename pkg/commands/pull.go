package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/pull"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "pull",
	Short:                 "Pull all repos that don't have any non-pushed changes",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/#-fgh-pull",
	Run: func(cmd *cobra.Command, args []string) {
		secrets, err := configuration.GetSecrets()
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		config, err := configuration.GetConfig(false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		clonedRepos, err := reposBasedOffCustomPath(cmd, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = pull.PullRepos(secrets, clonedRepos)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	// Allow the user to use this command on any directory
	err := addCustomPathFlag(pullCmd)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
