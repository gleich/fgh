package clean

import (
	"time"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/jedib0t/go-pretty/v6/progress"
)

// Get the repos that haven't been modified locally in a certain amount of time
func GetOutdated(
	pw progress.Writer,
	clonedRepos []repos.LocalRepo,
	yearsOld int,
	monthsOld int,
	daysOld int,
) ([]repos.DetailedLocalRepo, utils.CtxErr) {
	var (
		outdated      []repos.DetailedLocalRepo
		timeThreshold = time.Now().AddDate(-yearsOld, -monthsOld, -daysOld)
		formattedDate = utils.FormatDate(timeThreshold)
		tracker       = progress.Tracker{
			Message: "Checking if any repos were last updated before " + formattedDate,
			Total:   int64(len(clonedRepos)),
		}
	)
	tracker.SetValue(1)
	pw.AppendTracker(&tracker)

	for _, repo := range clonedRepos {
		notCommitted, notPushed, err := repos.WorkingState(repo.Path)
		if err.Error != nil {
			return []repos.DetailedLocalRepo{}, err
		}

		updatedTime, err := repos.LastUpdated(repo.Path)
		if err.Error != nil {
			return []repos.DetailedLocalRepo{}, err
		}

		if updatedTime.Unix() < timeThreshold.Unix() {
			outdated = append(outdated, repos.DetailedLocalRepo{
				Repo:         repo,
				ModTime:      updatedTime,
				NotCommitted: notCommitted,
				NotPushed:    notPushed,
			})
		}
		tracker.Increment(1)
	}

	return outdated, utils.CtxErr{}
}
