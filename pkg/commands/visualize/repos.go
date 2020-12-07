package visualize

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/briandowns/spinner"
)

// Get the repos for each user. User mapped to repo.
func GetRepos(clonedRepos []repos.LocalRepo) map[string][]repos.DetailedLocalRepo {
	mappedRepos := map[string][]repos.DetailedLocalRepo{}

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(" Getting data for %v repos", len(clonedRepos))
	spin.Start()

	for _, repo := range clonedRepos {
		var (
			notCommitted, notPushed = repos.WorkingState(repo.Path)
			updatedTime             = repos.LastUpdated(repo.Path)
		)
		mappedRepos[repo.Owner] = append(mappedRepos[repo.Owner], repos.DetailedLocalRepo{
			Repo:         repo,
			ModTime:      updatedTime,
			NotCommitted: notCommitted,
			NotPushed:    notPushed,
		})
	}

	spin.Stop()
	return mappedRepos
}
