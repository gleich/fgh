package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/migrate"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var mirgrateCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "migrate <FOLDER>",
	Short:                 "Migrate all the repos in a directory and its subdirectories",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/#-fgh-migrate",
	Run: func(cmd *cobra.Command, args []string) {
		folder, err := migrate.EnsureFolderExists(args)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		oldRepos, err := repos.Repos(folder, false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		config, err := configuration.GetConfig(false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		newPaths, err := migrate.NewPaths(oldRepos, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = migrate.ConfirmMove(newPaths)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = utils.MoveRepos(newPaths)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		_, err = clean.CleanUp(config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(mirgrateCmd)
}
