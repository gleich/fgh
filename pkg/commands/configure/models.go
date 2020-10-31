package configure

type RegularOutline struct {
	CloneClipboard bool `yaml:"clone_clipboard"`
}

type SecretsOutline struct {
	PAT      string `yaml:"pat"`
	Username string `yaml:"username"`
}
