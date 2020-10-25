package update

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/fatih/color"
)

// Ask the user if they want to move each repo
func AskMove(updated map[repos.LocalRepo]api.Repo) map[string]string {
	toMove := map[string]string{}
	for localRepo, repoAPIData := range updated {
		newPath := repos.RepoLocation(repoAPIData)
		fmt.Println()
		move := utils.Confirm(fmt.Sprintf(
			`Current Path: %v
  New Path:     %v
  Owner:        %v
  Name:         %v
  Language:     %v

  Would you like to move this repo to the new path?`,
			color.GreenString(localRepo.Path),
			color.GreenString(newPath),
			repoAPIData.Owner,
			repoAPIData.Name,
			repoAPIData.MainLanguage,
		))
		if move {
			toMove[localRepo.Path] = newPath
		}
	}

	if len(updated) != 0 {
		fmt.Println()
	}

	return toMove
}
