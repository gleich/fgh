package repos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterRepos(t *testing.T) {
	tt := []struct {
		repos          []LocalRepo
		expectedResult []LocalRepo
	}{
		{
			repos: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
			expectedResult: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "cli", Name: "cli"},
			},
			expectedResult: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "nuke"},
			},
			expectedResult: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
			},
		},
		{
			repos: []LocalRepo{
				{Owner: "Matt-Gleich", Name: "dots"},
				{Owner: "Matt-Gleich", Name: "dots"},
			},
			expectedResult: []LocalRepo{
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
