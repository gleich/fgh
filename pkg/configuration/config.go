package configuration

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

const (
	OwnerRep             = "OWNER"
	TypeRep              = "TYPE"
	LangRep              = "MAIN_LANGUAGE"
	DefaultStructureRoot = "github"
)

// Get the regular config configuration.
func GetConfig(ignoreErr bool) configure.RegularOutline {
	filePath := filepath.Join(configure.GetFolderPath(), configure.RegularFileName)
	var config configure.RegularOutline
	err := utils.ReadYAML(filePath, &config)
	if err.Error != nil && !ignoreErr {
		statuser.Error(err.Context, err.Error, 1)
	}

	// Setting defaults:
	if len(config.Structure) == 0 {
		config.Structure = []string{
			OwnerRep,
			TypeRep,
			LangRep,
		}
	}
	if config.StructureRoot == "" {
		config.StructureRoot = DefaultStructureRoot
	}
	if config.SpaceChar == "" {
		config.SpaceChar = "-"
	}

	return config
}
