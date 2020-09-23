package configure

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

func WriteConfiguration(config AnswersOutline) {
	folder := createFolders(true)
	fmt.Println(folder)
}

// Create the folder where the configuration should live
// Returns the folder path created
func createFolders(testing bool) string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get homedirectory", err, 1)
	}
	var folderPath string
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		folderPath = filepath.Join(homePath, ".config/fgh/")
	} else {
		folderPath = filepath.Join(homePath, ".fgh")
	}

	_, err = os.Stat(folderPath)
	if !os.IsNotExist(err) {
		override := utils.Confirm("Configuration already exists. Do you want to override it?")
		if !override {
			os.Exit(0)
		}
	}
	err = os.MkdirAll(folderPath, 0777)
	if err != nil {
		statuser.Error("Failed to create the configuration folder", err, 1)
	}
	return folderPath
}
