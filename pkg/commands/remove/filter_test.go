package remove

import (
	"testing"

	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/stretchr/testify/assert"
)

func TestFilterRepos(t *testing.T) {
	tt := []struct {
		repos          []repos.LocalRepo
		expectedResult []repos.LocalRepo
	}{
		{
			repos: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
			expectedResult: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "cli", Name: "cli"},
			},
			expectedResult: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "nuke"},
			},
			expectedResult: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "dots"},
			},
			expectedResult: []repos.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
	}

	for _, test := range tt {
		username := "Matt-Gleich"

		assert.Equal(t,
			test.expectedResult,
			FilterRepos(
				username,
				test.repos,
				[]string{"Matt-Gleich/dots"},
			),
		)

		assert.Equal(t,
			test.expectedResult,
			FilterRepos(
				username,
				test.repos,
				[]string{"dots"},
			),
		)

		assert.Equal(t,
			test.expectedResult,
			FilterRepos(
				username,
				test.repos,
				[]string{"Matt-Gleich/dots", ""},
			),
		)

	}
}
