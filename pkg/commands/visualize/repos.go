package visualize

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get the repos for each user. User mapped to repo.
func GetRepos(clonedRepos []repos.LocalRepo) map[string][]repos.DetailedLocalRepo {
	var numOfRepos int64
	for range clonedRepos {
		numOfRepos++
	}

	pw := utils.GenerateProgressWriter()
	tracker := progress.Tracker{
		Message: fmt.Sprintf("Getting information for %v repositories", numOfRepos),
		Total:   numOfRepos,
	}

	tracker.SetValue(1)
	pw.AppendTracker(&tracker)
	go pw.Render()

	mappedRepos := map[string][]repos.DetailedLocalRepo{}
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
		tracker.Increment(1)
	}

	return mappedRepos
}
