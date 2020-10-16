package clean

import (
	"os"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/statuser/v2"
)

func Remove(repos []location.LocalRepo) {
	for _, repo := range repos {
		err := os.RemoveAll(repo.Path)
		if err != nil {
			statuser.Error("Failed to remove "+repo.Path, err, 1)
		}
		statuser.Success("Removed " + repo.Path)
	}
}
