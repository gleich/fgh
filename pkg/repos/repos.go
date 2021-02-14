package repos

import (
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get all repos in the directory and all subdirectories
func Repos(rootPath string, ignoreErr bool) ([]LocalRepo, utils.CtxErr) {
	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = " Getting list of repos"
	spin.Start()

	var (
		oldRepos = []LocalRepo{}
		errCtx   utils.CtxErr
	)

	err := filepath.Walk(
		rootPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil && !ignoreErr {
				return err
			}

			isRepo, errInfo := IsGitRepo(path)
			if errInfo.Error != nil {
				errCtx = errInfo
				return errInfo.Error
			}

			if info.IsDir() && isRepo {
				absPath, err := filepath.Abs(path)
				if err != nil && !ignoreErr {
					errCtx = utils.CtxErr{
						Context: "Failed to get absolute path for " + path,
						Error:   err,
					}
					return err
				}

				owner, name, errInfo := OwnerAndNameFromRemote(path)
				if err != nil && !ignoreErr {
					errCtx = errInfo
					return errInfo.Error
				}

				oldRepos = append(oldRepos, LocalRepo{
					Owner: owner,
					Name:  name,
					Path:  absPath,
				})
			}
			return nil
		},
	)

	spin.Stop()
	if err != nil && !ignoreErr {
		return []LocalRepo{}, utils.CtxErr{
			Context: "Failed to get list of repos",
			Error:   err,
		}
	}
	if errCtx.Error != nil && !ignoreErr {
		return []LocalRepo{}, errCtx
	}

	if len(oldRepos) == 0 && !ignoreErr {
		statuser.Warning("0 repos found inside " + rootPath)
	}

	return oldRepos, utils.CtxErr{}
}
