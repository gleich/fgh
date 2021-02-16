package repos

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	gitignore "github.com/sabhiram/go-gitignore"
	"gopkg.in/djherbis/times.v1"
)

// Get the time the repo was lasted updated
func LastUpdated(path string) (time.Time, utils.CtxErr) {
	// Reading from gitignore and parsing it
	ignoreData, _ := ioutil.ReadFile(filepath.Join(path, ".gitignore"))
	ignore := gitignore.CompileIgnoreLines(string(ignoreData))

	var updatedTime time.Time
	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Checking if the file exists
			stat, err := os.Stat(path)
			if os.IsNotExist(err) || stat.Mode() == os.ModeSymlink {
				return nil
			}

			// Ignoring files in .gitignore
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			if ignore.MatchesPath(strings.TrimLeft(path, homeDir)) {
				return filepath.SkipDir
			}

			// Getting the modification time
			timeInfo, err := times.Stat(path)
			if err != nil {
				return err
			}
			modTime := timeInfo.ModTime()
			if modTime.Unix() > updatedTime.Unix() {
				updatedTime = modTime
			}
			return nil
		},
	)

	if err != nil {
		return time.Time{}, utils.CtxErr{
			Context: "Failed to get last updated time for " + path,
			Error:   err,
		}
	}

	return updatedTime, utils.CtxErr{}
}
