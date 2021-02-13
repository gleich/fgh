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
		folder := migrate.EnsureFolderExists(args)

		oldRepos, err := repos.Repos(folder, false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		config := configuration.GetConfig(false)
		newPaths := migrate.NewPaths(oldRepos, config)
		migrate.ConfirmMove(newPaths)
		utils.MoveRepos(newPaths)
		clean.CleanUp(config)
	},
}

func init() {
	rootCmd.AddCommand(mirgrateCmd)
}
