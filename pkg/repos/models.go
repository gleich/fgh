package repos

import "time"

// A repo already cloned locally
type LocalRepo struct {
	Owner string
	Name  string
	Path  string
}

// A more detailed sepc of a repo cloned locally
type DetailedLocalRepo struct {
	Repo         LocalRepo
	ModTime      time.Time
	NotCommitted bool
	NotPushed    bool
}
