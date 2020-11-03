package repos

import (
	"fmt"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/statuser/v2"
)

// Filter out repos that don't have the owner and name passed in via args
func FilterRepos(username string, repos []LocalRepo, args []string) (filtered []LocalRepo) {
	owner, name := clone.OwnerAndName(username, args)
	for _, repo := range repos {
		if strings.EqualFold(repo.Owner, owner) && strings.EqualFold(repo.Name, name) {
			filtered = append(filtered, repo)
		}
	}
	if len(filtered) == 0 {
		statuser.ErrorMsg(fmt.Sprintf("Failed to find %v/%v", owner, name), 1)
	}
	return filtered
}
