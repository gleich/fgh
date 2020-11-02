package migrate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get all repos in the directory and all subdirectories
func Repos(path string) []repos.LocalRepo {
	if !utils.HasInternetConnection() {
		statuser.Warning("Failed to establish an internet connection")
	}

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = " Getting list of repos"
	spin.Start()

	oldRepos := []repos.LocalRepo{}
	err := filepath.Walk(
		path,
		func(cwd string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && repos.IsGitRepo(cwd) {
				owner, name := repos.OwnerAndNameFromRemote(cwd)
				oldRepos = append(oldRepos, repos.LocalRepo{
					Owner: owner,
					Name:  name,
					Path:  cwd,
				})
			}
			return nil
		},
	)

	spin.Stop()
	if err != nil {
		statuser.Error("Failed to get list of repos", err, 1)
	}

	if len(oldRepos) == 0 {
		statuser.Warning("0 repos found inside " + path)
	}

	statuser.Success(fmt.Sprintf("Detected %v repos in %v", len(oldRepos), path))
	return oldRepos
}
