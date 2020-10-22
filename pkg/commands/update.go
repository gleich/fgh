package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/update"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for any repos that have updated and need to go to a different folder. Won't move anything without your confirmation",
	Run: func(cmd *cobra.Command, args []string) {
		repos := location.Repos()
		changedRepos := update.GetChanged(repos)
		toMove := update.ConfirmMove(changedRepos)
		update.MoveRepos(toMove)
		clean.CleanUp()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
