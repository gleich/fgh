package repos

import (
	"path/filepath"
	"testing"

	"github.com/gleich/fgh/pkg/api"
	"github.com/gleich/fgh/pkg/commands/configure"
	"github.com/gleich/fgh/pkg/configuration"
	"github.com/stretchr/testify/assert"
)

func TestRepoLocation(t *testing.T) {
	result, err := RepoLocation(
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
	assert.NoError(t, err.Error)
	root, err := StructureRootPath(configure.RegularOutline{
		StructureRoot: "github/",
	})
	assert.NoError(t, err.Error)
	assert.Equal(t, filepath.Join(root, "/Testing-Owner/private/Go/fgh"), result)

	result, err = RepoLocation(
		api.Repo{
			Owner:        "gleich",
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
	assert.NoError(t, err.Error)
	root, err = StructureRootPath(configure.RegularOutline{
		StructureRoot: "/code/stuff",
	})
	assert.NoError(t, err.Error)
	assert.Equal(t, filepath.Join(root, "/gleich/JavaScript/archived/site-v2"), result)

	result, err = RepoLocation(
		api.Repo{
			Owner:        "gleich",
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
	assert.NoError(t, err.Error)
	root, err = StructureRootPath(configure.RegularOutline{
		StructureRoot: "/code/stuff",
	})
	assert.NoError(t, err.Error)
	assert.Equal(t, filepath.Join(root, "/gleich/JavaScript/site-v2"), result)
}
