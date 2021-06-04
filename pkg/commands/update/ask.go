package update

import (
	"fmt"

	"github.com/gleich/fgh/pkg/api"
	"github.com/gleich/fgh/pkg/commands/configure"
	"github.com/gleich/fgh/pkg/repos"
	"github.com/gleich/fgh/pkg/utils"
	"github.com/fatih/color"
)

// Ask the user if they want to move each repo
func AskMove(
	updated map[repos.LocalRepo]api.Repo,
	config configure.RegularOutline,
) (map[string]string, utils.CtxErr) {
	toMove := map[string]string{}
	for localRepo, repoAPIData := range updated {
		newPath, err := repos.RepoLocation(repoAPIData, config)
		if err.Error != nil {
			return map[string]string{}, err
		}

		fmt.Println()
		move, err := utils.Confirm(fmt.Sprintf(
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

		if err.Error != nil {
			return map[string]string{}, err
		}

		if move {
			toMove[localRepo.Path] = newPath
		}
	}

	if len(updated) != 0 {
		fmt.Println()
	}

	return toMove, utils.CtxErr{}
}
