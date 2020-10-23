package remove

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/statuser/v2"
)

// Filter out repos that don't have the owner and name passed in via args
func FilterRepos(secrets configure.SecretsOutline, repos []location.LocalRepo, args []string) (filtered []location.LocalRepo) {
	owner, name := clone.OwnerAndName(secrets, args)
	for _, repo := range repos {
		if (repo.Owner == owner) && (repo.Name == name) {
			filtered = append(filtered, repo)
		}
	}
	if len(filtered) == 0 {
		statuser.ErrorMsg("Failed to find repo", 1)
	}
	return filtered
}
