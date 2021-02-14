package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "clean",
	Short:                 "Ask to remove old or deleted cloned repos",
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-clean",
	Run: func(cmd *cobra.Command, args []string) {
		flags, err := clean.ParseFlags(cmd)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		config, err := configuration.GetConfig(false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		clonedRepos, err := reposBasedOffCustomPath(cmd, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		var (
			toRemove    = []repos.LocalRepo{}
			progressBar = utils.GenerateProgressWriter()
		)

		go progressBar.Render()

		if !flags.SkipOutdated {
			outdated, err := clean.GetOutdated(progressBar, clonedRepos, flags.Years, flags.Months, flags.Days)
			if err.Error != nil {
				statuser.Error(err.Context, err.Error, 1)
			}

			approved, err := clean.AskToRemoveOutdated(outdated)
			if err.Error != nil {
				statuser.Error(err.Context, err.Error, 1)
			}

			toRemove = append(toRemove, approved...)
		}

		if !flags.SkipDeleted {
			deleted, err := clean.GetDeleted(progressBar, clonedRepos)
			if err.Error != nil {
				statuser.Error(err.Context, err.Error, 1)
			}

			approved, err := clean.AskToRemoveDeleted(deleted)
			if err.Error != nil {
				statuser.Error(err.Context, err.Error, 1)
			}

			toRemove = append(toRemove, approved...)
		}

		err = clean.Remove(toRemove)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		_, err = clean.CleanUp(config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().IntP("years", "y", 0, "Counts as outdated repo if it isn't modified locally in given number of years (default 0)")
	cleanCmd.Flags().IntP("months", "m", 2, "Counts as outdated repo if it isn't modified locally in given number of months")
	cleanCmd.Flags().IntP("days", "d", 0, "Counts as outdated repo if it isn't modified locally in given number of days (default 0)")
	cleanCmd.Flags().Bool("skipOutdated", false, "Don't check for outdated repos")
	cleanCmd.Flags().Bool("skipDeleted", false, "Don't check for deleted repos")

	// Allow the user to use this command on any directory
	err := addCustomPathFlag(visualizeCmd)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
