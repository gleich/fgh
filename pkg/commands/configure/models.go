package configure

type RegularOutline struct {
	CloneClipboard    bool     `yaml:"clone_clipboard"`
	Structure         []string `yaml:"structure"`
	StructureRoot     string   `yaml:"structure_root"`
	DontAppendHomeDir bool     `yaml:"dont_append_home_dir"`
	LowercaseLang     bool     `yaml:"lowercase_lang"`
	SpaceChar         string   `yaml:"space_character"`
}

type SecretsOutline struct {
	PAT      string `yaml:"pat"`
	Username string `yaml:"username"`
}
