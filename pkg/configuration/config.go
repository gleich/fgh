package configuration

import (
	"os"
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

const (
	OwnerRep             = "REPO_OWNER"
	TypeRep              = "REPO_TYPE"
	LangRep              = "REPO_MAIN_LANGUAGE"
	DefaultStructureRoot = "github"
)

// Get the regular config configuration
func GetConfig() configure.RegularOutline {
	filePath := filepath.Join(configure.GetFolderPath(), configure.RegularFileName)
	var config configure.RegularOutline
	err := utils.ReadYAML(filePath, &config)
	if err != nil {
		statuser.Error("Failed to read from configuration", err, 1)
	}

	// Setting defaults:
	if len(config.Structure) == 0 {
		config.Structure = []string{
			OwnerRep,
			TypeRep,
			LangRep,
		}
	}
	if config.StructureRoot == "" { // If not defined in the config; default
		homePath, err := os.UserHomeDir()
		if err != nil {
			statuser.Error("Failed to get home directory", err, 1)
		}
		config.StructureRoot = filepath.Join(homePath, DefaultStructureRoot)
	}

	return config
}
