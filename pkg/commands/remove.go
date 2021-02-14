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

		filtered, err := repos.FilterRepos(secrets.Username, clonedRepos, args)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = remove.RemoveRepos(filtered)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		_, err = clean.CleanUp(config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
	},
	ValidArgsFunction: reposAsValidArgs,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Allow the user to use this command on any directory
	err := addCustomPathFlag(visualizeCmd)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
