package clean

import (
	"fmt"
	"strings"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	tf "github.com/hepsiburada/time-formatter"
)

// Confirm with the user that they want to remove an outdated repo
func AskToRemoveOutdated(outdatedRepos []OutdatedRepo) []location.LocalRepo {
	toRemove := []location.LocalRepo{}
	for _, repo := range outdatedRepos {
		time := formatDate(repo.ModTime)
		remove := utils.Confirm(fmt.Sprintf(
			`Path:     %v
  Owner:    %v
  Name:     %v
  Language: %v
  Type:     %v
  Last Local Update: %v

  Would you like to remove this repo?`,
			repo.Repo.Path,
			repo.Repo.Owner,
			repo.Repo.Name,
			repo.Repo.Language,
			strings.Title(repo.Repo.Type),
			color.GreenString(time),
		))
		fmt.Println()
		if remove {
			toRemove = append(toRemove, repo.Repo)
		}
	}
	return toRemove
}

// Confirm with the user that they want to remove a certain
func AskToRemoveDeleted(deletedRepos []location.LocalRepo) []location.LocalRepo {
	if len(deletedRepos) != 0 {
		fmt.Println("\n----------------------")
		fmt.Println(" Deleted Repositories ")
		fmt.Println("----------------------")
	}

	toRemove := []location.LocalRepo{}
	for _, repo := range deletedRepos {
		remove := utils.Confirm(fmt.Sprintf(
			"It seems as though %v has been deleted on GitHub.\n  Would you like to remove it locally?",
			color.GreenString(repo.Path),
		))
		fmt.Println()
		if remove {
			toRemove = append(toRemove, repo)
		}
	}
	return toRemove
}

// Format date in the following format:
// December 25th, 2020 at 12:00PM
func formatDate(date time.Time) string {
	formatter := tf.New()
	return formatter.To(date, fmt.Sprintf(
		"%s %s, %s at %v:%v%s",
		tf.MMMM,
		humanize.Ordinal(date.Day()),
		tf.YYYY,
		date.Hour(),
		date.Minute(),
		tf.A,
	))
}
