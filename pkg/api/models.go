package api

type Repo struct {
	Owner        string
	Name         string
	MainLanguage string
	Private      bool
	Archived     bool
	Template     bool
	Disabled     bool
	Mirror       bool
	Fork         bool
}
