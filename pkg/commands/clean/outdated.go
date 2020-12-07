package clean

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get the repos that haven't been modified locally in a certain amount of time
func GetOutdated(clonedRepos []repos.LocalRepo, yearsOld int, monthsOld int, daysOld int) (outdated []repos.DetailedLocalRepo) {
	timeThreshold := time.Now().AddDate(-yearsOld, -monthsOld, -daysOld)
	formattedDate := utils.FormatDate(timeThreshold)

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(
		" Checking for any repos last updated before %v",
		formattedDate,
	)
	spin.Start()

	for _, repo := range clonedRepos {
		var (
			notCommitted, notPushed = repos.WorkingState(repo.Path)
			updatedTime             = repos.LastUpdated(repo.Path)
		)
		if updatedTime.Unix() < timeThreshold.Unix() {
			outdated = append(outdated, repos.DetailedLocalRepo{
				Repo:        repo,
				ModTime:     updatedTime,
				Uncommitted: notCommitted,
				NotPushed:   notPushed,
			})
		}
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("%v repos last updated locally before %v", len(outdated), formattedDate))
	return outdated
}
