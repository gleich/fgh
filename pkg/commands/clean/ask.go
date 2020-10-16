package clean

import (
	"fmt"
	"strings"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/dustin/go-humanize"
	tf "github.com/hepsiburada/time-formatter"
)

func AskToRemove(outdatedRepos []OutdatedRepo) (toRemove []location.LocalRepo) {
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
			time,
		))
		fmt.Println()
		if remove {
			toRemove = append(toRemove, repo.Repo)
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
