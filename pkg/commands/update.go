package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/update"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "update",
	Short:                 "Ask if you want to update the path of any repos with updated fields",
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-update",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			config       = configuration.GetConfig(false)
			repos        = repos.ReposInStructure(config)
			changedRepos = update.GetChanged(repos, config)
			toMove       = update.AskMove(changedRepos, config)
		)
		utils.MoveRepos(toMove)
		clean.CleanUp(config)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
