package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
)

// Move all repos
func MoveRepos(repos map[string]string) CtxErr {
	for oldPath, newPath := range repos {
		// Making base folder
		parts := strings.Split(newPath, string(filepath.Separator))
		baseFolder := strings.Join(parts[:len(parts)-1], string(filepath.Separator))
		err := os.MkdirAll(baseFolder, 0777)
		if err != nil {
			return CtxErr{
				Context: "Failed to make " + baseFolder,
				Error:   err,
			}
		}

		// Renaming folder to new path
		err = os.Rename(oldPath, newPath)
		if err != nil {
			return CtxErr{
				Error:   err,
				Context: "Failed to move " + oldPath + " to " + newPath,
			}
		}
		statuser.Success("Moved " + oldPath + " to " + newPath)
	}
	return CtxErr{}
}
