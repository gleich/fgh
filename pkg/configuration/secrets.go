package configuration

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
)

func Secrets() configure.SecretsOutline {
	folderPath := configure.GetFolderPath()
	filePath := filepath.Join(folderPath, configure.SecretsFileName)
	utils.ReadYAML(filePath)
}
