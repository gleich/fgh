package clean

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/djherbis/times.v1"
)

type OutdatedRepo struct {
	Repo    location.LocalRepo
	ModTime time.Time
}

// Get the repos that haven't been modified locally in a certain amount of time
func Outdated(repos []location.LocalRepo, yearsOld int, monthsOld int, daysOld int) (outdated []OutdatedRepo) {
	timeThreshold := time.Now().AddDate(-yearsOld, -monthsOld, -daysOld)
	for _, repo := range repos {
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
			statuser.Error("Failed to get updated time for "+repo.Path, err, 1)
		}
		if updatedTime.Unix() < timeThreshold.Unix() {
			outdated = append(outdated, OutdatedRepo{Repo: repo, ModTime: updatedTime})
		}
	}
	statuser.Success(fmt.Sprintf("Got %v repos last updated before %v", len(outdated), formatDate(timeThreshold)))
	return outdated
}
