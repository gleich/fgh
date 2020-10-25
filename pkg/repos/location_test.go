package repos

import (
	"path/filepath"
	"testing"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestRepoLocation(t *testing.T) {
	result1 := RepoLocation(api.Repo{
		Owner:        "Testing-Owner",
		Name:         "fgh",
		MainLanguage: "Go",
		Private:      true,
	})
	assert.Equal(t, filepath.Join(GitHubFolder(), "/Testing-Owner/private/Go/fgh"), result1)

	result2 := RepoLocation(api.Repo{
		Owner:        "Matt-Gleich",
		Name:         "site-v2",
		MainLanguage: "JavaScript",
		Archived:     true,
		Private:      true,
	})
	assert.Equal(t, filepath.Join(GitHubFolder(), "/Matt-Gleich/archived/JavaScript/site-v2"), result2)
}
