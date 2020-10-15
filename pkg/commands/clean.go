package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

// TODO:
// Fix outdated as it gets the modified time for the folder and not the contents

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Ask user if they want to remove old repos",
	Run: func(cmd *cobra.Command, args []string) {
		flags := clean.ParseFlags(cmd)
		repos := location.Repos()
		outdated := clean.Outdated(repos, flags.Years, flags.Months, flags.Days)
		clean.AskToRemove(outdated)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().IntP("years", "y", 0, "Counts as outdated repo if it isn't modified locally in given number of years (default 0)")
	cleanCmd.Flags().IntP("months", "m", 2, "Counts as outdated repo if it isn't modified locally in given number of months")
	cleanCmd.Flags().IntP("days", "d", 0, "Counts as outdated repo if it isn't modified locally in given number of days (default 0)")
}
