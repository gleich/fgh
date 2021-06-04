package visualize

import (
	"sort"

	"github.com/gleich/fgh/pkg/repos"
	"github.com/gleich/fgh/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get the repos for each user. User mapped to repo.
func GetRepos(clonedRepos []repos.LocalRepo) (map[string][]repos.DetailedLocalRepo, utils.CtxErr) {
	var numOfRepos int64
	for range clonedRepos {
		numOfRepos++
	}

	pw := utils.GenerateProgressWriter()
	tracker := progress.Tracker{
		Message: "Getting information for repositories",
		Total:   numOfRepos,
	}

	tracker.SetValue(1)
	pw.AppendTracker(&tracker)
	go pw.Render()

	detailedRepos := []repos.DetailedLocalRepo{}
	for _, repo := range clonedRepos {
		notCommitted, notPushed, err := repos.WorkingState(repo.Path)
		if err.Error != nil {
			return map[string][]repos.DetailedLocalRepo{}, err
		}

		updatedTime, err := repos.LastUpdated(repo.Path)
		if err.Error != nil {
			return map[string][]repos.DetailedLocalRepo{}, err
		}
		detailedRepos = append(detailedRepos, repos.DetailedLocalRepo{
			Repo:         repo,
			ModTime:      updatedTime,
			NotCommitted: notCommitted,
			NotPushed:    notPushed,
		})
		tracker.Increment(1)
	}

	// Sorting the repos by mod time
	sort.SliceStable(detailedRepos, func(i, j int) bool {
		return detailedRepos[i].ModTime.After(detailedRepos[j].ModTime)
	})

	// Grouping the repos by name
	groupedRepos := map[string][]repos.DetailedLocalRepo{}
	for _, repo := range detailedRepos {
		groupedRepos[repo.Repo.Owner] = append(groupedRepos[repo.Repo.Owner], repo)
	}

	return groupedRepos, utils.CtxErr{}
}
