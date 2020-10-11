package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "update",
	Short: "Remove all repos not updated after a certain amount of time",
	Run: func(cmd *cobra.Command, args []string) {
		flags := clean.ParseFlags(cmd)
		repos := location.Repos()
		outdated := clean.Outdated(repos, flags.Years, flags.Months, flags.Days)
		fmt.Println(outdated)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().IntP("years", "y", 0, "Counts as outdated repo if it isn't modified locally in given number of years (default 0)")
	cleanCmd.Flags().IntP("months", "m", 2, "Counts as outdated repo if it isn't modified locally in given number of months")
	cleanCmd.Flags().IntP("days", "d", 0, "Counts as outdated repo if it isn't modified locally in given number of days (default 0)")
}
