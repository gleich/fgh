package repos

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/djherbis/times.v1"
)

// Get the time the repo was lasted updated
func LastUpdated(path string) time.Time {
	var updatedTime time.Time
	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() && info.Name() == "node_modules" {
				return filepath.SkipDir
			}

			info, err := times.Stat(path)
			if err != nil {
				return err
			}
			modTime := info.ModTime()
			if modTime.Unix() > updatedTime.Unix() {
				updatedTime = modTime
			}
			return nil
		},
	)
	if err != nil {
		statuser.Error("Failed to get last updated time for "+path, err, 1)
	}
	return updatedTime
}
