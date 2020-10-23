package clean

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"github.com/enescakir/emoji"
)

// Get all the repos locally that have been deleted on GitHub
func GetDeleted(repos []location.LocalRepo) (deleted []location.LocalRepo) {
	if !utils.HasInternetConnection() {
		statuser.Warning("Failed to establish an internet connection")
	}

	spin := spinner.New(utils.SpinnerCharSet, 40*time.Millisecond)
	spin.Suffix = fmt.Sprintf("  %v  Checking if any local repos have been deleted from GitHub", emoji.Information)
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
