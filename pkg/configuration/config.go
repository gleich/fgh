package configuration

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
)

const (
	OwnerRep             = "OWNER"
	TypeRep              = "TYPE"
	LangRep              = "MAIN_LANGUAGE"
	DefaultStructureRoot = "github"
)

// Get the regular config configuration.
func GetConfig(ignoreErr bool) (configure.RegularOutline, utils.CtxErr) {
	configPath, err := configure.GetFolderPath()
	if err.Error != nil {
		return configure.RegularOutline{}, err
	}

	filePath := filepath.Join(configPath, configure.RegularFileName)
	var config configure.RegularOutline

	err = utils.ReadYAML(filePath, &config)
	if err.Error != nil && !ignoreErr {
		return configure.RegularOutline{}, err
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

	return config, utils.CtxErr{}
}
