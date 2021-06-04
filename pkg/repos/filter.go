package repos

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gleich/fgh/pkg/commands/clone"
	"github.com/gleich/fgh/pkg/utils"
)

// Filter out repos that don't have the owner and name passed in via args
func FilterRepos(username string, repos []LocalRepo, args []string) ([]LocalRepo, utils.CtxErr) {
	var filtered []LocalRepo

	owner, name, err := clone.OwnerAndName(username, args)
	if err.Error != nil {
		return []LocalRepo{}, err
	}

	for _, repo := range repos {
		if strings.EqualFold(repo.Name, name) {
			filtered = append(filtered, repo)
		}
	}

	// Give repos owned by user greater precedence
	sort.Slice(filtered, func(i int, j int) bool {
		if filtered[j].Owner == owner && filtered[i].Owner != owner {
			return false
		}

		return true
	})

	if len(filtered) == 0 {
		errMsg := fmt.Sprintf("Failed to find %v/%v", owner, name)
		return []LocalRepo{}, utils.CtxErr{
			Context: errMsg,
			Error:   errors.New(errMsg),
		}
	}

	return filtered, utils.CtxErr{}
}
