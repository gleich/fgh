package clean

import (
	"github.com/gleich/fgh/pkg/api"
	"github.com/gleich/fgh/pkg/configuration"
	"github.com/gleich/fgh/pkg/repos"
	"github.com/gleich/fgh/pkg/utils"
	"github.com/gleich/statuser/v2"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get all the repos locally that have been deleted on GitHub
func GetDeleted(
	pw progress.Writer,
	clonedRepos []repos.LocalRepo,
) ([]repos.LocalRepo, utils.CtxErr) {
	if !utils.HasInternetConnection() {
		statuser.Warning("Failed to establish an internet connection")
	}

	secrets, err := configuration.GetSecrets()
	if err.Error != nil {
		return []repos.LocalRepo{}, err
	}

	var (
		client  = api.GenerateClient(secrets.PAT)
		tracker = progress.Tracker{
			Message: "Checking if any repos have been deleted from GitHub",
			Total:   int64(len(clonedRepos)),
		}
	)
	tracker.SetValue(1)
	pw.AppendTracker(&tracker)

	var deleted []repos.LocalRepo
	for _, localRepo := range clonedRepos {
		_, err := api.RepoData(client, localRepo.Owner, localRepo.Name)
		if err.Error != nil {
			deleted = append(deleted, localRepo)
		}
		tracker.Increment(1)
	}

	return deleted, utils.CtxErr{}
}
