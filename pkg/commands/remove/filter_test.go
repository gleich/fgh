package remove

import (
	"testing"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/stretchr/testify/assert"
)

func TestFilterRepos(t *testing.T) {
	tt := []struct {
		repos          []location.LocalRepo
		expectedResult []location.LocalRepo
	}{
		{
			repos: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
			expectedResult: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "cli", Name: "cli"},
			},
			expectedResult: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "nuke"},
			},
			expectedResult: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "dots"},
			},
			expectedResult: []location.LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
	}

	for _, test := range tt {
		secrets := configure.SecretsOutline{Username: "Matt-Gleich"}

		assert.Equal(t,
			test.expectedResult,
			FilterRepos(
				secrets,
				test.repos,
				[]string{"Matt-Gleich/dots"},
			),
		)

		assert.Equal(t,
			test.expectedResult,
			FilterRepos(
				secrets,
				test.repos,
				[]string{"dots"},
			),
		)

		assert.Equal(t,
			test.expectedResult,
			FilterRepos(
				secrets,
				test.repos,
				[]string{"Matt-Gleich/dots", ""},
			),
		)

	}
}
