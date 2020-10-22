package remove

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

// Ask to remove each repo and then remove it
func RemoveRepos(repos []location.LocalRepo) {
	for _, repo := range repos {
		remove := utils.Confirm(fmt.Sprintf(
			"Are you sure you want to permanently remove %v from your computer?", repo.Path,
		))
		if remove {
			err := os.RemoveAll(repo.Path)
			if err != nil {
				statuser.Error("Failed to remove "+repo.Path, err, 1)
			}
			statuser.Success("Removed " + repo.Path)
		}
	}
}
