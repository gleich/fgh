package repos

import (
	"path/filepath"
	"testing"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/stretchr/testify/assert"
)

func TestRepoLocation(t *testing.T) {
	result1 := RepoLocation(
		api.Repo{
			Owner:        "Testing-Owner",
			Name:         "fgh",
			MainLanguage: "Go",
			Archived:     false,
			Private:      true,
		},
		configure.RegularOutline{
			Structure: []string{
				configuration.OwnerRep,
				configuration.TypeRep,
				configuration.LangRep,
			},
			StructureRoot: "github/",
		},
	)
	assert.Equal(t, filepath.Join(StructureRootFolder(configure.RegularOutline{
		StructureRoot: "github/",
	}), "/Testing-Owner/private/Go/fgh"), result1)

	result2 := RepoLocation(
		api.Repo{
			Owner:        "Matt-Gleich",
			Name:         "site-v2",
			MainLanguage: "JavaScript",
			Archived:     true,
			Private:      true,
		},
		configure.RegularOutline{
			Structure: []string{
				configuration.OwnerRep,
				configuration.LangRep,
				configuration.TypeRep,
			},
			StructureRoot: "/code/stuff",
		},
	)
	assert.Equal(t, filepath.Join(StructureRootFolder(configure.RegularOutline{
		StructureRoot: "/code/stuff",
	}), "/Matt-Gleich/JavaScript/archived/site-v2"), result2)

	result3 := RepoLocation(
		api.Repo{
			Owner:        "Matt-Gleich",
			Name:         "site-v2",
			MainLanguage: "JavaScript",
			Archived:     false,
			Private:      false,
		},
		configure.RegularOutline{
			Structure: []string{
				configuration.OwnerRep,
				configuration.LangRep,
			},
			StructureRoot: "/code/stuff",
		},
	)
	assert.Equal(t, filepath.Join(StructureRootFolder(configure.RegularOutline{
		StructureRoot: "/code/stuff",
	}), "/Matt-Gleich/JavaScript/site-v2"), result3)
}
