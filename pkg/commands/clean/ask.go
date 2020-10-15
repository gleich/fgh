package clean

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/dustin/go-humanize"
	tf "github.com/hepsiburada/time-formatter"
)

func AskToRemove(outdatedRepos []OutdatedRepo) (toRemove []location.LocalRepo) {
	for _, repo := range outdatedRepos {
		formatter := tf.New()
		time := formatter.To(repo.ModTime, fmt.Sprintf(
			"%s %s %s",
			tf.MMMM,
			humanize.Ordinal(repo.ModTime.Day()),
			tf.YYYY,
		))
		remove := utils.Confirm(fmt.Sprintf(
			"%v/%v hasn't been updated since %v\n  Would you like to remove it?",
			repo.Repo.Owner,
			repo.Repo.Name,
			time,
		))
		if remove {
			toRemove = append(toRemove, repo.Repo)
		}
	}
	return toRemove
}
