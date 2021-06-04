package commands

import (
	"github.com/gleich/fgh/pkg/commands/clean"
	"github.com/gleich/fgh/pkg/commands/remove"
	"github.com/gleich/fgh/pkg/configuration"
	"github.com/gleich/fgh/pkg/repos"
	"github.com/gleich/fgh/pkg/utils"
	"github.com/gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "remove <OWNER/NAME>",
	Short:                 "Remove a cloned repo",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		force, err := utils.GetBool("force", cmd)

		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		secrets, err := configuration.GetSecrets()
		if err.Error != nil && !force {
			statuser.Error(err.Context, err.Error, 1)
		}

		config, err := configuration.GetConfig(false)
		if err.Error != nil && !force {
			statuser.Error(err.Context, err.Error, 1)

		}

		clonedRepos, err := reposBasedOffCustomPath(cmd, config)
		if err.Error != nil && !force {
			statuser.Error(err.Context, err.Error, 1)
		}

		filtered, err := repos.FilterRepos(secrets.Username, clonedRepos, args)
		if err.Error != nil && !force {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = remove.RemoveRepos(filtered, force)
		if err.Error != nil && !force {
			statuser.Error(err.Context, err.Error, 1)
		}

		_, err = clean.CleanUp(config)
		if err.Error != nil && !force {
			statuser.Error(err.Context, err.Error, 1)
		}
	},
	ValidArgsFunction: reposAsValidArgs,
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().Bool("force", false, "Force remove even if there are some errors")

	// Allow the user to use this command on any directory
	err := addCustomPathFlag(removeCmd)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
