package clean

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/Matt-Gleich/statuser/v2"
)

// Remove the repos
func Remove(repos []location.LocalRepo) {
	for _, repo := range repos {
		err := os.RemoveAll(repo.Path)
		if err != nil {
			statuser.Error("Failed to remove "+repo.Path, err, 1)
		}
		statuser.Success("Removed " + repo.Path)
	}
}

// Remove all fgh folders that don't have anything in them
func CleanUp() (removed []string) {
	ghFolder := location.GitHubFolder()
	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			trimmedPath := strings.TrimPrefix(path, ghFolder)
			parts := strings.Split(trimmedPath, string(filepath.Separator))
			if len(parts) > 4 {
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
		statuser.Error("Failed to get list of folders in github folder located at root", err, 1)
	}

	for _, folder := range removed {
		err = os.Remove(folder)
		if err != nil {
			statuser.Error("Failed to remove "+folder, err, 1)
		}
	}

	return removed
}
