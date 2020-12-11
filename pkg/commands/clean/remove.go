package clean

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

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

// Remove empty folders in the structure (NOT EMPTY REPOS)
func CleanUp(config configure.RegularOutline) []string {
	ghFolder := repos.StructureRootPath(config)

	foldersToCheck := []string{}
	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				if repos.IsGitRepo(path) {
					return filepath.SkipDir
				}
				foldersToCheck = append(foldersToCheck, path)
			}
			return nil
		},
	)
	if err != nil {
		statuser.Error("Failed to get list of folders in structure", err, 1)
	}

	// Sorting paths by length
	sort.Slice(foldersToCheck, func(i, j int) bool {
		return len(foldersToCheck[j]) < len(foldersToCheck[i])
	})

	for _, path := range foldersToCheck {
		content, err := ioutil.ReadDir(path)
		if err != nil {
			statuser.Error("Failed to get contents of "+path, err, 1)
		}
		if len(content) == 0 {
			err = os.Remove(path)
			if err != nil {
				statuser.Error("Failed to remove "+path, err, 1)
			}
		} else if len(content) == 1 && !content[0].IsDir() && content[0].Name() == ".DS_Store" {
			// If the folder only contains a .DS_Store file
			err = os.RemoveAll(path)
			if err != nil {
				statuser.Error("Failed to remove "+path, err, 1)
			}
		}
	}
	return foldersToCheck
}
