package clean

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get all the repos locally that have been deleted on GitHub
func GetDeleted(repos []repos.LocalRepo) (deleted []repos.LocalRepo) {
	if !utils.HasInternetConnection() {
		statuser.Warning("Failed to establish an internet connection")
	}

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = " Checking if any local repos have been deleted from GitHub"
	spin.Start()

	client := api.GenerateClient(configuration.GetSecrets().PAT)
	for _, localRepo := range repos {
		_, err := api.RepoData(client, localRepo.Owner, localRepo.Name)
		if err != nil {
			deleted = append(deleted, localRepo)
		}
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("%v repos that are deleted from GitHub and cloned locally", len(deleted)))
	return deleted
}
