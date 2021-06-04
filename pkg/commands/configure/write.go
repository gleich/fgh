package configure

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/gleich/fgh/pkg/utils"
)

const (
	RegularFileName = "config.yaml"
	SecretsFileName = "secrets.yaml"
)

// Write the regular configuration for the program
func WriteConfig(config RegularOutline) utils.CtxErr {
	folderPath, err := GetFolderPath()
	if err.Error != nil {
		return err
	}

	filePath := filepath.Join(folderPath, RegularFileName)
	err = utils.WriteYAML(config, filePath)
	if err.Error != nil {
		return err
	}
	return utils.CtxErr{}
}

// Write the secret configuration for the program
func WriteSecrets(secrets SecretsOutline) utils.CtxErr {
	folderPath, err := GetFolderPath()
	if err.Error != nil {
		return err
	}

	filePath := filepath.Join(folderPath, SecretsFileName)
	err = utils.WriteYAML(secrets, filePath)
	if err.Error != nil {
		return err
	}
	return utils.CtxErr{}
}

// Get the folder path for where the configuration should live
func GetFolderPath() (string, utils.CtxErr) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", utils.CtxErr{
			Context: "Failed to get homedirectory",
			Error:   err,
		}
	}

	var folderPath string
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		folderPath = filepath.Join(homePath, ".config", "fgh")
	} else {
		folderPath = filepath.Join(homePath, ".fgh")
	}

	return folderPath, utils.CtxErr{}
}
