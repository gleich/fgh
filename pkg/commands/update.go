package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/update"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Ask if you want to update the path of any repos that should have a new path",
	Long:  "If a repo changes its type, main language, owner, or name the path to your local repo won't match. Running fgh update will check every repo to see if it should have a new path. If it does it will ask you if you want to move the entire repo to that new path. So if I had this repo cloned and then I was to archive it the path would change from ~/github/Matt-Gleich/public/Go/fgh/ to ~/github/Matt-Gleich/archived/Go/fgh/.",
	Run: func(cmd *cobra.Command, args []string) {
		repos := location.Repos()
		changedRepos := update.GetChanged(repos)
		toMove := update.AskMove(changedRepos)
		update.MoveRepos(toMove)
		clean.CleanUp()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
