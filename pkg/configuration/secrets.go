package configuration

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/commands/login"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
)

// Get the secret configuration
func GetSecrets() configure.SecretsOutline {
	filePath := filepath.Join(configure.GetFolderPath(), configure.SecretsFileName)
	var config configure.SecretsOutline
	err := utils.ReadYAML(filePath, &config)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
	if config.Username == "" {
		username := login.Username(config.PAT)
		config = configure.SecretsOutline{PAT: config.PAT, Username: username}
		configure.WriteSecrets(config)
	}
	return config
}
