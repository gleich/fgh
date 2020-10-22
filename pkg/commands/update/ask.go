package update

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/fatih/color"
)

func ConfirmMove(updated map[location.LocalRepo]api.Repo) map[string]string {
	toMove := map[string]string{}
	for localRepo, repoAPIData := range updated {
		newPath := location.RepoLocation(repoAPIData)
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
	fmt.Println()
	return toMove
}
