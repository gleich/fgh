package repos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/statuser/v2"
)

// Get the root GitHub folder
func StructureRootFolder(structureRoot string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home directory", err, 1)
	}
	return filepath.Join(homeDir, structureRoot)
}

// Get the location to clone the repo
func RepoLocation(repo api.Repo, config configure.RegularOutline) string {
	path := filepath.Join(StructureRootFolder(config.StructureRoot), filepath.Join(config.Structure...))

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

	return filepath.Join(path, repo.Name)
}
