package configuration

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

// Get the regular config configuration
func GetConfig() configure.RegularOutline {
	folderPath := configure.GetFolderPath()
	filePath := filepath.Join(folderPath, configure.RegularFileName)
	var config configure.RegularOutline
	err := utils.ReadYAML(filePath, &config)
	if err != nil {
		statuser.Error("Failed to read from configuration", err, 1)
	}
	return config
}
