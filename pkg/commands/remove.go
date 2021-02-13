package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/remove"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "remove <OWNER/NAME>",
	Short:                 "Remove a cloned repo",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			secrets     = configuration.GetSecrets()
			config      = configuration.GetConfig(false)
			clonedRepos = reposBasedOffCustomPath(cmd, config)
		)

		filtered, err := repos.FilterRepos(secrets.Username, clonedRepos, args)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		remove.RemoveRepos(filtered)
		clean.CleanUp(config)
	},
	ValidArgsFunction: reposAsValidArgs,
}

func init() {
	rootCmd.AddCommand(removeCmd)
	addCustomPathFlag(removeCmd)
}
