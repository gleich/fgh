package update

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
)

// Move all repos
func MoveRepos(repos map[string]string) {
	for oldPath, newPath := range repos {
		// Making folder
		parts := strings.Split(newPath, string(filepath.Separator))
		baseFolder := strings.Join(parts[:len(parts)-1], string(filepath.Separator))
		err := os.MkdirAll(baseFolder, 0777)
		if err != nil {
			statuser.Error("Failed to make "+baseFolder, err, 1)
		}

		err = os.Rename(oldPath, newPath)
		if err != nil {
			statuser.Error("Failed to move "+oldPath+" to "+newPath, err, 1)
		}
		statuser.Success("Moved " + oldPath + " to " + newPath)
	}
}
