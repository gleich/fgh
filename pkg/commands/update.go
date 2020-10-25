package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/update"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("%v  Ask if you want to update the path of any repos with updated fields", emoji.UpArrow),
	Long:  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-update",
	Run: func(cmd *cobra.Command, args []string) {
		repos := repos.Repos()
		changedRepos := update.GetChanged(repos)
		toMove := update.AskMove(changedRepos)
		update.MoveRepos(toMove)
		clean.CleanUp()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
