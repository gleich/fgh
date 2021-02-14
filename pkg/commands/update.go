package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/update"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "update",
	Short:                 "Ask if you want to update the path of any repos with updated fields",
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-update",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configuration.GetConfig(false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		repos, err := repos.ReposInStructure(config, false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		changedRepos, err := update.GetChanged(repos, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		toMove, err := update.AskMove(changedRepos, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = utils.MoveRepos(toMove)
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
	rootCmd.AddCommand(updateCmd)
}
