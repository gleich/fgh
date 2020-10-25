package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: fmt.Sprintf("%v Ask to remove old or deleted cloned repos", emoji.Soap),
	Long:  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-clean",
	Run: func(cmd *cobra.Command, args []string) {
		flags := clean.ParseFlags(cmd)
		clonedRepos := repos.Repos()

		outdated := clean.GetOutdated(clonedRepos, flags.Years, flags.Months, flags.Days)
		toRemove := clean.AskToRemoveOutdated(outdated)

		deleted := clean.GetDeleted(clonedRepos)
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
