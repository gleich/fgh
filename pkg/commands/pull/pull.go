package pull

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
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
			statuser.Warning(fmt.Sprintf("%v/%v has %v\n", repo.Owner, repo.Name, statusMsg))
			continue
		}

		err := os.Chdir(repo.Path)
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to change directory to " + repo.Path,
				Error:   err,
			}
		}

		output, err := exec.Command("git", "pull").Output()
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to pull changes for " + repo.Path,
				Error:   err,
			}
		}
		if strings.Contains(string(output), "Already up to date.") {
			fmt.Printf("%v/%v is already up to date\n", repo.Owner, repo.Name)
			continue
		}

		statuser.Success(fmt.Sprintf("Pulled latest changes for %v/%v", repo.Owner, repo.Name))
		pulled++
	}
	fmt.Println()
	statuser.Success(fmt.Sprintf("Pulled latest changes for %v repos", pulled))

	return utils.CtxErr{}
}
