package configure

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

const (
	RegularFileName = "config.yaml"
	SecretsFileName = "secrets.yaml"
)

// Write the regular configuration for the program
func WriteConfig(config RegularOutline) {
	folderPath := GetFolderPath()
	filePath := filepath.Join(folderPath, RegularFileName)
	err := utils.WriteYAML(config, filePath)
	if err != nil {
		statuser.Error("Failed to write config secrets", err, 1)
	}
}

// Write the secret configuration for the program
func WriteSecrets(secrets SecretsOutline) {
	folderPath := GetFolderPath()
	filePath := filepath.Join(folderPath, SecretsFileName)
	err := utils.WriteYAML(secrets, filePath)
	if err != nil {
		statuser.Error("Failed to write config secrets", err, 1)
	}
}

// Get the folder path for where the configuration should live
func GetFolderPath() string {
	homePath, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get homedirectory", err, 1)
	}
	var folderPath string
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		folderPath = filepath.Join(homePath, ".config", "fgh")
	} else {
		folderPath = filepath.Join(homePath, ".fgh")
	}
	return folderPath
}
