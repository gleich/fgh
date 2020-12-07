package clean

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/fatih/color"
)

// Confirm with the user that they want to remove an outdated repo
func AskToRemoveOutdated(outdatedRepos []repos.DetailedLocalRepo) []repos.LocalRepo {
	toRemove := []repos.LocalRepo{}
	for _, repo := range outdatedRepos {
		time := utils.FormatDate(repo.ModTime)

		var (
			uncommittedMsg = color.GreenString("None")
			notPushedMsg   = color.GreenString("None")
		)
		if !repo.Uncommitted {
			uncommittedMsg = color.RedString("Yes")
		}
		if !repo.NotPushed {
			notPushedMsg = color.RedString("Yes")
		}

		remove := utils.Confirm(fmt.Sprintf(
			`Path:     %v
  Owner:    %v
  Name:     %v
  Last Local Update:   %v
  Uncommitted changes: %v
  Changes not pushed:  %v

  Would you like to remove this repo locally?`,
			repo.Repo.Path,
			repo.Repo.Owner,
			repo.Repo.Name,
			color.GreenString(time),
			uncommittedMsg,
			notPushedMsg,
		))
		fmt.Println()
		if remove {
			toRemove = append(toRemove, repo.Repo)
		}
	}
	return toRemove
}

// Confirm with the user that they want to remove a deleted repo
func AskToRemoveDeleted(deletedRepos []repos.LocalRepo) []repos.LocalRepo {
	if len(deletedRepos) != 0 {
		fmt.Println("\n----------------------")
		fmt.Println(" Deleted Repositories ")
		fmt.Println("----------------------")
	}

	toRemove := []repos.LocalRepo{}
	for _, repo := range deletedRepos {
		remove := utils.Confirm(fmt.Sprintf(
			"It seems as though %v has been deleted on GitHub.\n  Would you like to remove it?",
			color.GreenString(repo.Path),
		))
		fmt.Println()
		if remove {
			toRemove = append(toRemove, repo)
		}
	}
	return toRemove
}
