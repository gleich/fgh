package configure

type RegularOutline struct {
	CloneClipboard bool     `yaml:"clone_clipboard"`
	Structure      []string `yaml:"structure"`
	StructureRoot  string   `yaml:"structure_root"`
}

type SecretsOutline struct {
	PAT      string `yaml:"pat"`
	Username string `yaml:"username"`
}
