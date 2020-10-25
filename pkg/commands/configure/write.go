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

// Write the configuration
func WriteConfiguration(config RegularOutline) {
	configFolder := CreateFolders()
	WriteConfig(configFolder, config)
	statuser.Success("Wrote to config")
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

// Create the folder where the configuration should live
// Returns the folder path created
func CreateFolders() string {
	folderPath := GetFolderPath()
	_, err := os.Stat(folderPath)
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

// Write the regular configuration for the program
func WriteConfig(folder string, config RegularOutline) {
	filePath := filepath.Join(folder, RegularFileName)
	err := utils.WriteYAML(config, filePath)
	if err != nil {
		statuser.Error("Failed to write config secrets", err, 1)
	}
}

// Write the secret configuration for the program
func WriteSecrets(folder string, secrets SecretsOutline) {
	filePath := filepath.Join(folder, SecretsFileName)
	err := utils.WriteYAML(secrets, filePath)
	if err != nil {
		statuser.Error("Failed to write config secrets", err, 1)
	}
}
