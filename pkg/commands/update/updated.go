package update

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"github.com/enescakir/emoji"
	"github.com/fatih/color"
)

// Get all repos cloned locally that have a new location based off the repo changes
func GetChanged(repos []location.LocalRepo) map[location.LocalRepo]api.Repo {
	spin := spinner.New(spinner.CharSets[1], 40*time.Millisecond)
	spin.Suffix = fmt.Sprintf(" %v  Getting latest repo information for all cloned all repos", emoji.Information)
	spin.Start()

	updated := map[location.LocalRepo]api.Repo{}
	client := api.GenerateClient()
	for _, localRepo := range repos {
		updatedData, err := api.RepoData(client, localRepo.Owner, localRepo.Name)
		if err != nil {
			statuser.Error(
				fmt.Sprintf(
					"Failed to get data about repo located in %v. It is possible that it got deleted. If you want to remove all deleted repos run the clean command.",
					color.RedString(localRepo.Path),
				),
				err, 1,
			)
		}
		if location.RepoLocation(updatedData) != localRepo.Path {
			updated[localRepo] = updatedData
		}
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("%v repos cloned locally with new paths", len(updated)))
	return updated
}
