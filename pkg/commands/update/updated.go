package update

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Get all repos cloned locally that have a new location based off the repo changes
func GetChanged(clonedRepos []repos.LocalRepo) map[repos.LocalRepo]api.Repo {
	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = " Getting latest repo information for all cloned all repos"
	spin.Start()

	updated := map[repos.LocalRepo]api.Repo{}
	client := api.GenerateClient(configuration.GetSecrets().PAT)
	for _, localRepo := range clonedRepos {
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
		if repos.RepoLocation(updatedData) != localRepo.Path {
			updated[localRepo] = updatedData
		}
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("%v repos cloned locally with new paths", len(updated)))
	return updated
}
