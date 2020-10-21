package clean

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get all the repos locally that have been deleted
func GetDeleted(repos []location.LocalRepo) (deleted []location.LocalRepo) {
	if !utils.HasInternetConnection() {
		statuser.Warning("Failed to establish an internet connection")
	}

	spin := spinner.New(spinner.CharSets[1], 40*time.Millisecond)
	spin.Suffix = " ℹ️  Checking if any local repos have been deleted from GitHub"
	spin.Start()

	for _, localRepo := range repos {
		_, err := api.RepoData(api.GenerateClient(), localRepo.Owner, localRepo.Name)
		if err != nil {
			deleted = append(deleted, localRepo)
		}
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("%v repos that are deleted from GitHub and cloned locally", len(deleted)))
	return deleted
}
