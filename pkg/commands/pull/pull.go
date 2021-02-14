package pull

import (
	"errors"
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Pull the latest changes for all local repos not in
// a working state
func PullRepos(secrets configure.SecretsOutline, clonedRepos []repos.LocalRepo) utils.CtxErr {
	var pulled int
	for _, repo := range clonedRepos {
		// Outputting the status message
		var (
			statusMsg   string
			pullChanges bool
		)

		committed, pushed, errCtx := repos.WorkingState(repo.Path)
		if errCtx.Error != nil {
			return errCtx
		}

		if !committed && !pushed {
			statusMsg = "changes not committed and not pushed"
		} else if !committed {
			statusMsg = "changes not committed"
		} else if !pushed {
			statusMsg = "changes not pushed"
		} else {
			pullChanges = true
		}
		if !pullChanges {
			msg := fmt.Sprintf("%v/%v has %v\n", repo.Owner, repo.Name, statusMsg)
			return utils.CtxErr{
				Context: msg,
				Error:   errors.New(msg),
			}
		}

		// Pulling latest changes
		gitRepo, err := git.PlainOpen(repo.Path)
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to open repo in: " + repo.Path,
				Error:   err,
			}
		}

		workingTree, err := gitRepo.Worktree()
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to get working tree for " + repo.Path,
				Error:   err,
			}
		}

		err = workingTree.Pull(&git.PullOptions{
			RemoteName: "origin",
			Auth: &http.BasicAuth{
				Username: secrets.Username,
				Password: secrets.PAT,
			},
		})

		// Outputting final message
		if err != nil {
			if err.Error() == "already up-to-date" {
				fmt.Printf("%v/%v is already up to date\n", repo.Owner, repo.Name)
				continue
			}
			return utils.CtxErr{
				Context: "Failed to pull changes for " + repo.Path,
				Error:   err,
			}
		}
		statuser.Success(fmt.Sprintf("Pulled latest changes for %v/%v", repo.Owner, repo.Name))
		pulled++
	}
	fmt.Println()
	statuser.Success(fmt.Sprintf("Pulled latest changes for %v repos", pulled))

	return utils.CtxErr{}
}
