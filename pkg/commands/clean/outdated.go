package clean

import (
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
	twoMonthsAgo := time.Now().AddDate(-yearsOld, -monthsOld, -daysOld)
	for _, repo := range repos {
		t, err := times.Stat(repo.Path)
		if err != nil {
			statuser.Error("Failed to get modified time for "+repo.Path, err, 1)
		}
		mod := t.ModTime()
		if mod.Unix() < twoMonthsAgo.Unix() {
			outdated = append(outdated, OutdatedRepo{
				Repo:    repo,
				ModTime: mod,
			})
		}
	}
	return outdated
}
