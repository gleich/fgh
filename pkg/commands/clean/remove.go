package clean

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/statuser/v2"
)

// Remove the repos
func Remove(repos []repos.LocalRepo) {
	for _, repo := range repos {
		err := os.RemoveAll(repo.Path)
		if err != nil {
			statuser.Error("Failed to remove "+repo.Path, err, 1)
		}
		statuser.Success("Removed " + repo.Path)
	}
}

// Remove empty folders >3 diretories deep from ~/github
func CleanUp(config configure.RegularOutline) (removed []string) {
	ghFolder := repos.GitHubFolder(config.StructureRoot)
	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			trimmedPath := strings.TrimPrefix(path, ghFolder)
			parts := strings.Split(trimmedPath, string(filepath.Separator))
			if len(parts) > len(config.Structure)+2 {
				return filepath.SkipDir
			} else if info.IsDir() {
				f, err := ioutil.ReadDir(path)
				if err != nil {
					statuser.Error("Failed to list directory: "+path, err, 1)
				}
				if len(f) == 0 {
					removed = append(removed, path)
				}
			}
			return nil
		},
	)
	if err != nil {
		statuser.Error("Failed to remove empty folders >3 diretories deep from "+config.StructureRoot, err, 1)
	}

	for _, folder := range removed {
		err = os.Remove(folder)
		if err != nil {
			statuser.Error("Failed to remove "+folder, err, 1)
		}
	}

	return removed
}
