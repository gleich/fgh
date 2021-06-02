package update

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get all repos cloned locally that have a new location based off the repo changes
func GetChanged(
	clonedRepos []repos.LocalRepo,
	config configure.RegularOutline,
) (map[repos.LocalRepo]api.Repo, utils.CtxErr) {
	updated := map[repos.LocalRepo]api.Repo{}
	secrets, err := configuration.GetSecrets()
	if err.Error != nil {
		return map[repos.LocalRepo]api.Repo{}, err
	}

	var (
		client  = api.GenerateClient(secrets.PAT)
		pw      = utils.GenerateProgressWriter()
		tracker = progress.Tracker{
			Message: "Getting valid paths for all repos",
			Total:   int64(len(clonedRepos)),
		}
	)

	tracker.SetValue(1)
	pw.AppendTracker(&tracker)
	go pw.Render()

	for _, localRepo := range clonedRepos {
		updatedData, err := api.RepoData(client, localRepo.Owner, localRepo.Name)
		if err.Error != nil {
			return map[repos.LocalRepo]api.Repo{}, utils.CtxErr{
				Context: fmt.Sprintf(
					"Failed to get data about repo located in %v. It is possible that it got deleted. If you want to remove all deleted repos run the clean command.",
					localRepo.Path,
				),
				Error: err.Error,
			}
		}

		location, err := repos.RepoLocation(updatedData, config)
		if err.Error != nil {
			return map[repos.LocalRepo]api.Repo{}, err
		}

		if location != localRepo.Path {
			updated[localRepo] = updatedData
		}
		tracker.Increment(1)
	}

	return updated, utils.CtxErr{}
}
