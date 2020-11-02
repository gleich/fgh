package migrate

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/statuser/v2"
)

// Check to make sure that the folder exist
// Returns the path of the folder
func EnsureFolderExists(args []string) string {
	folder, err := rawEnsureFolderExists(args)
	if err != nil {
		statuser.Error("Invalid folder", err, 1)
	}
	return folder
}

// Testable, core logic for EnsureFolderExists
func rawEnsureFolderExists(args []string) (string, error) {
	folder := args[0]
	state, err := os.Stat(folder)
	if os.IsNotExist(err) {
		return folder, fmt.Errorf("%v doesn't exist", folder)
	}
	if !state.IsDir() {
		return folder, fmt.Errorf("%v isn't a folder", folder)
	}
	return folder, nil
}
