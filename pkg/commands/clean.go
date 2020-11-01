package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "clean",
	Short:                 fmt.Sprintf("%v Ask to remove old or deleted cloned repos", emoji.Soap),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-clean",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			flags       = clean.ParseFlags(cmd)
			config      = configuration.GetConfig()
			clonedRepos = repos.Repos(config)
			toRemove    = []repos.LocalRepo{}
		)

		if !flags.SkipOutdated {
			outdated := clean.GetOutdated(clonedRepos, flags.Years, flags.Months, flags.Days)
			toRemove = append(toRemove, clean.AskToRemoveOutdated(outdated)...)
		}

		if !flags.SkipDeleted {
			deleted := clean.GetDeleted(clonedRepos)
			toRemove = append(toRemove, clean.AskToRemoveDeleted(deleted)...)
		}

		clean.Remove(toRemove)
		clean.CleanUp(config)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().IntP("years", "y", 0, "Counts as outdated repo if it isn't modified locally in given number of years (default 0)")
	cleanCmd.Flags().IntP("months", "m", 2, "Counts as outdated repo if it isn't modified locally in given number of months")
	cleanCmd.Flags().IntP("days", "d", 0, "Counts as outdated repo if it isn't modified locally in given number of days (default 0)")
	cleanCmd.Flags().Bool("skipOutdated", false, "Don't check for outdated repos")
	cleanCmd.Flags().Bool("skipDeleted", false, "Don't check for deleted repos")
}
