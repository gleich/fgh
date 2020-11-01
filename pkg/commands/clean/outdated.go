package clean

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"gopkg.in/djherbis/times.v1"
)

type OutdatedRepo struct {
	Repo        repos.LocalRepo
	ModTime     time.Time
	Uncommitted bool
	NotPushed   bool
}

// Get the repos that haven't been modified locally in a certain amount of time
func GetOutdated(clonedRepos []repos.LocalRepo, yearsOld int, monthsOld int, daysOld int) (outdated []OutdatedRepo) {
	timeThreshold := time.Now().AddDate(-yearsOld, -monthsOld, -daysOld)
	formattedDate := formatDate(timeThreshold)

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(
		" Checking for any repos last updated before %v",
		formattedDate,
	)
	spin.Start()

	for _, repo := range clonedRepos {
		var updatedTime time.Time
		err := filepath.Walk(
			repo.Path,
			func(path string, _ os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				info, err := times.Stat(path)
				if err != nil {
					return err
				}
				modTime := info.ModTime()
				if modTime.Unix() > updatedTime.Unix() {
					updatedTime = modTime
				}
				return nil
			},
		)
		if err != nil {
			statuser.Error("Failed to get last updated time for "+repo.Path, err, 1)
		}

		committed, pushed := repos.WorkingState(repo.Path)
		if updatedTime.Unix() < timeThreshold.Unix() {
			outdated = append(outdated, OutdatedRepo{
				Repo:        repo,
				ModTime:     updatedTime,
				Uncommitted: committed,
				NotPushed:   pushed,
			})
		}
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("%v repos last updated locally before %v", len(outdated), formattedDate))
	return outdated
}
