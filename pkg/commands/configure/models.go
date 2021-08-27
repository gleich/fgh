package configure

type RegularOutline struct {
	CloneClipboard bool     `yaml:"clone_clipboard"`
	Structure      []string `yaml:"structure"`
	StructureRoot  string   `yaml:"structure_root"`
	LowercaseLang  bool     `yaml:"lowercase_lang"`
	SpaceChar      string   `yaml:"space_character"`
	SSH            bool     `yaml:"ssh"`
}

type SecretsOutline struct {
	PAT      string `yaml:"pat"`
	Username string `yaml:"username"`
}
