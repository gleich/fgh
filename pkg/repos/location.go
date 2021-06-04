package repos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gleich/fgh/pkg/api"
	"github.com/gleich/fgh/pkg/commands/configure"
	"github.com/gleich/fgh/pkg/configuration"
	"github.com/gleich/fgh/pkg/utils"
)

// Get the structure root path
func StructureRootPath(config configure.RegularOutline) (string, utils.CtxErr) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", utils.CtxErr{
			Context: "Failed to get home directory",
			Error:   err,
		}
	}

	root := config.StructureRoot
	if strings.HasPrefix(root, string(filepath.Separator)) {
		return root, utils.CtxErr{}
	} else {
		return filepath.Join(homeDir, root), utils.CtxErr{}
	}
}

// Get the location to clone the repo
func RepoLocation(repo api.Repo, config configure.RegularOutline) (string, utils.CtxErr) {
	root, err := StructureRootPath(config)
	if err.Error != nil {
		return "", err
	}

	path := filepath.Join(root, filepath.Join(config.Structure...))

	// Replacing owner
	path = strings.ReplaceAll(path, configuration.OwnerRep, repo.Owner)

	// Replacing type
	if repo.Template {
		path = strings.ReplaceAll(path, configuration.TypeRep, "template")
	} else if repo.Disabled {
		path = strings.ReplaceAll(path, configuration.TypeRep, "disabled")
	} else if repo.Archived {
		path = strings.ReplaceAll(path, configuration.TypeRep, "archived")
	} else if repo.Mirror {
		path = strings.ReplaceAll(path, configuration.TypeRep, "mirror")
	} else if repo.Fork {
		path = strings.ReplaceAll(path, configuration.TypeRep, "fork")
	} else if repo.Private {
		path = strings.ReplaceAll(path, configuration.TypeRep, "private")
	} else {
		path = strings.ReplaceAll(path, configuration.TypeRep, "public")
	}

	// Replacing lang
	if config.LowercaseLang {
		repo.MainLanguage = strings.ToLower(repo.MainLanguage)
	}
	repo.MainLanguage = strings.ReplaceAll(repo.MainLanguage, " ", config.SpaceChar)
	path = strings.ReplaceAll(path, configuration.LangRep, repo.MainLanguage)

	return filepath.Join(path, repo.Name), utils.CtxErr{}
}
