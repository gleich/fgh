package migrate

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get the new paths for all the repos
func NewPaths(
	oldRepos []repos.LocalRepo,
	config configure.RegularOutline,
) (map[string]string, utils.CtxErr) {
	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(" Getting latest metadata for %v repos", len(oldRepos))
	spin.Start()

	newPaths := map[string]string{}
	secrets, err := configuration.GetSecrets()
	if err.Error != nil {
		return map[string]string{}, err
	}

	client := api.GenerateClient(secrets.PAT)
	for _, repo := range oldRepos {
		metadata, err := api.RepoData(client, repo.Owner, repo.Name)
		if err.Error != nil {
			statuser.Warning(fmt.Sprintf(
				"%v will not be moved because it has either been deleted from github or you don't have access",
				repo.Path,
			))
		}

		newLocation, err := repos.RepoLocation(metadata, config)
		if err.Error != nil {
			return map[string]string{}, err
		}
		newPaths[repo.Path] = newLocation
	}
	spin.Stop()

	if len(newPaths) == 0 {
		os.Exit(0)
	}

	statuser.Success(fmt.Sprintf("Got latest metadata for %v repos", len(oldRepos)))
	return newPaths, utils.CtxErr{}
}
