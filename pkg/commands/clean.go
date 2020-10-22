package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Ask if you want to remove repos not updated in awhile or were deleted from GitHub.",
	Long: `When you run this command fgh will check every single repo for two things:

1. If it hasn't been modified locally in a certain amount of time. The default amount of time is 2 months but this can be changed with flags. See fgh clean --help for more info.
2. If the repo has been deleted on GitHub.

If either of those conditions are met fgh will ask you if you would like to remove it and shows you some information about the repo. This only removes the repo locally.`,
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
