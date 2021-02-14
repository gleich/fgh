package repos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
)

// Get all cloned repos in fgh's file structure
func ReposInStructure(config configure.RegularOutline, ignoreErr bool) ([]LocalRepo, utils.CtxErr) {
	var repos []LocalRepo

	ghFolder, errCtx := StructureRootPath(config)
	if errCtx.Error != nil {
		return []LocalRepo{}, errCtx
	}

	err := filepath.Walk(
		ghFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil && !ignoreErr {
				return err
			}

			parts := strings.Split((strings.TrimPrefix(path, ghFolder)), string(filepath.Separator))
			if len(parts) > len(config.Structure)+2 {
				return filepath.SkipDir
			}

			isRepo, errInfo := IsGitRepo(path)
			if errInfo.Error != nil {
				errCtx = errInfo
			}

			if len(parts) == len(config.Structure)+2 && info.IsDir() && isRepo {
				owner, name, err := OwnerAndNameFromRemote(path)
				if err.Error != nil && !ignoreErr {
					errCtx = err
					return err.Error
				}

				repos = append(repos, LocalRepo{
					Owner: owner,
					Name:  name,
					Path:  path,
				})
			}
			return nil
		},
	)

	if err != nil && !ignoreErr {
		return []LocalRepo{}, utils.CtxErr{
			Context: "Failed to get list cloned of repos",
			Error:   err,
		}
	}

	if errCtx.Error != nil && !ignoreErr {
		return []LocalRepo{}, errCtx
	}

	return repos, utils.CtxErr{}
}
