package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Ask user if they want to remove repos not updated in a while or were deleted from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		flags := clean.ParseFlags(cmd)
		repos := location.Repos()

		outdated := clean.GetOutdated(repos, flags.Years, flags.Months, flags.Days)
		toRemove := clean.AskToRemoveOutdated(outdated)

		deleted := clean.GetDeleted(repos)
		toRemove = append(toRemove, clean.AskToRemoveDeleted(deleted)...)

		clean.Remove(toRemove)
		clean.CleanUp()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().IntP("years", "y", 0, "Counts as outdated repo if it isn't modified locally in given number of years (default 0)")
	cleanCmd.Flags().IntP("months", "m", 2, "Counts as outdated repo if it isn't modified locally in given number of months")
	cleanCmd.Flags().IntP("days", "d", 0, "Counts as outdated repo if it isn't modified locally in given number of days (default 0)")
}
