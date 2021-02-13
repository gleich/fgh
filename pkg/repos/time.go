package repos

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"gopkg.in/djherbis/times.v1"
)

// Get the time the repo was lasted updated
func LastUpdated(path string) (time.Time, utils.CtxErr) {
	var updatedTime time.Time
	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Checking if the file exists
			_, err = os.Stat(path)
			if os.IsNotExist(err) {
				return nil
			}

			// Ignoring node_modules
			if info.IsDir() && info.Name() == "node_modules" {
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
