package repos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
)

// A repo already cloned locally
type LocalRepo struct {
	Owner string
	Name  string
	Path  string
}

// Get all cloned repos in fgh's file structure
func ReposInStructure(config configure.RegularOutline) (repos []LocalRepo) {
	ghFolder := GitHubFolder(config.StructureRoot)

	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			parts := strings.Split((strings.TrimPrefix(path, ghFolder)), string(filepath.Separator))
			if len(parts) > len(config.Structure)+2 {
				return filepath.SkipDir
			}

			if len(parts) == len(config.Structure)+2 && info.IsDir() && IsGitRepo(path) {
				owner, name := OwnerAndNameFromRemote(path)
				repos = append(repos, LocalRepo{
					Owner: owner,
					Name:  name,
					Path:  path,
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
