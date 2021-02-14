package clean

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

// Remove the repos
func Remove(repos []repos.LocalRepo) utils.CtxErr {
	for _, repo := range repos {
		err := os.RemoveAll(repo.Path)
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to remove " + repo.Path,
				Error:   err,
			}
		}
		statuser.Success("Removed " + repo.Path)
	}
	return utils.CtxErr{}
}

// Remove empty folders in the structure (NOT EMPTY REPOS)
func CleanUp(config configure.RegularOutline) ([]string, utils.CtxErr) {
	ghFolder, errCtx := repos.StructureRootPath(config)
	if errCtx.Error != nil {
		return []string{}, errCtx
	}

	foldersToCheck := []string{}
	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				isRepo, errCtx := repos.IsGitRepo(path)
				if errCtx.Error != nil {
					return errCtx.Error
				}
				if isRepo {
					return filepath.SkipDir
				}
				foldersToCheck = append(foldersToCheck, path)
			}
			return nil
		},
	)
	if err != nil {
		return []string{}, utils.CtxErr{
			Context: "Failed to get list of folders in structure",
			Error:   err,
		}
	}

	// Sorting paths by length
	sort.Slice(foldersToCheck, func(i, j int) bool {
		return len(foldersToCheck[j]) < len(foldersToCheck[i])
	})

	for _, path := range foldersToCheck {
		content, err := ioutil.ReadDir(path)
		if err != nil {
			return []string{}, utils.CtxErr{
				Context: "Failed to get contents of " + path,
				Error:   err,
			}
		}
		if len(content) == 0 {
			err = os.Remove(path)
			return []string{}, utils.CtxErr{
				Context: "Failed to remove " + path,
				Error:   err,
			}
		} else if len(content) == 1 && !content[0].IsDir() && content[0].Name() == ".DS_Store" {
			// If the folder only contains a .DS_Store file
			err = os.RemoveAll(path)
			if err != nil {
				return []string{}, utils.CtxErr{
					Context: "Failed to remove " + path,
					Error:   err,
				}
			}
		}
	}
	return foldersToCheck, utils.CtxErr{}
}
