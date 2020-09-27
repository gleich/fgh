package clone

import (
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/statuser/v2"
)

// Get the location to clone the repo
func Location(repo Repository) string {
	var path string
	path, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home directory", err, 1)
	}
	path = filepath.Join(path, "github", repo.Owner)
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
