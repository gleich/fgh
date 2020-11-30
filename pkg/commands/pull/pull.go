package pull

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Pull the latest changes for all local repos not in
// a working state
func PullRepos(secrets configure.SecretsOutline, clonedRepos []repos.LocalRepo) {
	var pulled int
	for _, repo := range clonedRepos {
		// Outputting the status message
		var (
			statusMsg         string
			pullChanges       bool
			committed, pushed = repos.WorkingState(repo.Path)
		)
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
			statuser.Warning(fmt.Sprintf("%v/%v has %v\n", repo.Owner, repo.Name, statusMsg))
			continue
		}

		// Pulling latest changes
		gitRepo, err := git.PlainOpen(repo.Path)
		if err != nil {
			statuser.Error("Failed to open repo in: "+repo.Path, err, 1)
		}
		workingTree, err := gitRepo.Worktree()
		if err != nil {
			statuser.Error("Failed to get working tree for "+repo.Path, err, 1)
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
			statuser.Error("Failed to pull changes for "+repo.Path, err, 1)
		}
		statuser.Success(fmt.Sprintf("Pulled latest changes for %v/%v", repo.Owner, repo.Name))
		pulled++
	}
	fmt.Println()
	statuser.Success(fmt.Sprintf("Pulled latest changes for %v repos", pulled))
}
