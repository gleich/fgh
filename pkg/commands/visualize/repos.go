package visualize

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get the repos for each user. User mapped to repo.
func GetRepos(clonedRepos []repos.LocalRepo) map[string][]repos.DetailedLocalRepo {
	mappedRepos := map[string][]repos.DetailedLocalRepo{}

	var numOfRepos int
	for range clonedRepos {
		numOfRepos++
	}
	numOfRepos--

	pw := progress.NewWriter()
	pw.SetTrackerLength(30)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.SetUpdateFrequency(utils.SpinnerSpeed)
	pw.Style().Colors = progress.StyleCircle.Colors
	pw.SetStyle(progress.StyleCircle)

	tracker := progress.Tracker{
		Message: fmt.Sprintf("Getting information for %v repositoties", numOfRepos),
		Total:   int64(numOfRepos),
		Units:   progress.UnitsDefault,
	}
	pw.AppendTracker(&tracker)
	go pw.Render()

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
