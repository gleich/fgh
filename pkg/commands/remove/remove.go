package remove

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

// Ask to remove each repo and then remove it
func RemoveRepos(clonedRepos []repos.LocalRepo, force bool) utils.CtxErr {
	for _, repo := range clonedRepos {
		committed, pushed, errCtx := repos.WorkingState(repo.Path)
		if errCtx.Error != nil && !force {
			return errCtx
		}

		if !committed {
			statuser.Warning(
				fmt.Sprintf("Repository located at %v has uncommitted changes", repo.Path),
			)
		}
		if !pushed {
			statuser.Warning(
				fmt.Sprintf("Repository located at %v has changes not pushed to a remote", repo.Path),
			)
		}
		remove, errCtx := utils.Confirm(fmt.Sprintf(
			"Are you sure you want to permanently remove %v from your computer?", repo.Path,
		))
		if errCtx.Error != nil {
			return errCtx
		}
		if !remove {
			continue
		}

		err := os.RemoveAll(repo.Path)
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to remove " + repo.Path,
				Error:   err,
			}
		}
		statuser.Success("Removed " + repo.Path)
	}
	return utils.CtxErr{}
}
