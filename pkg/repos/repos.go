package repos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
)

// A repo already cloned locally
type LocalRepo struct {
	Owner    string
	Name     string
	Type     string
	Language string
	Path     string
}

// Get the root GitHub folder
func GitHubFolder() string {
	var path string
	path, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home directory", err, 1)
	}
	return filepath.Join(path, "github")
}

// Get all cloned repos in fgh's file structure
func Repos() (repos []LocalRepo) {
	ghFolder := GitHubFolder()

	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			parts := strings.Split((strings.TrimPrefix(path, ghFolder)), string(filepath.Separator))
			if len(parts) > 5 {
				return filepath.SkipDir
			}

			if len(parts) == 5 && info.IsDir() && isGitRepo(path) {
				owner, name := OwnerAndNameFromRemote(path)
				repos = append(repos, LocalRepo{
					Owner:    owner,
					Name:     name,
					Type:     parts[2],
					Language: parts[3],
					Path:     path,
				})
			}
			return nil
		},
	)

	if err != nil {
		statuser.Error("Failed to get list cloned of repos", err, 1)
	}
	return repos
}
