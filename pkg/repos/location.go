package repos

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/api"
)

// Get the location to clone the repo
func RepoLocation(repo api.Repo) string {
	path := filepath.Join(GitHubFolder(), repo.Owner)
	if repo.Template {
		path = filepath.Join(path, "template")
	} else if repo.Disabled {
		path = filepath.Join(path, "disabled")
	} else if repo.Archived {
		path = filepath.Join(path, "archived")
	} else if repo.Mirror {
		path = filepath.Join(path, "mirror")
	} else if repo.Fork {
		path = filepath.Join(path, "fork")
	} else if repo.Private {
		path = filepath.Join(path, "private")
	} else {
		path = filepath.Join(path, "public")
	}
	path = filepath.Join(path, repo.MainLanguage, repo.Name)
	return path
}
