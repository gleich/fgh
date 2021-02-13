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
		config := configuration.GetConfig(false)

		repos, err := repos.ReposInStructure(config, false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		changedRepos := update.GetChanged(repos, config)
		toMove, err := update.AskMove(changedRepos, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		utils.MoveRepos(toMove)
		clean.CleanUp(config)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
