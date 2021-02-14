package configuration

import (
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/commands/login"
	"github.com/Matt-Gleich/fgh/pkg/utils"
)

// Get the secret configuration
func GetSecrets() (configure.SecretsOutline, utils.CtxErr) {
	configPath, err := configure.GetFolderPath()
	if err.Error != nil {
		return configure.SecretsOutline{}, err
	}

	filePath := filepath.Join(configPath, configure.SecretsFileName)
	var config configure.SecretsOutline
	err = utils.ReadYAML(filePath, &config)
	if err.Error != nil {
		return configure.SecretsOutline{}, err
	}
	if config.Username == "" {
		username, err := login.Username(config.PAT)
		if err.Error != nil {
			return configure.SecretsOutline{}, err
		}
		config = configure.SecretsOutline{PAT: config.PAT, Username: username}
		configure.WriteSecrets(config)
	}
	return config, utils.CtxErr{}
}
