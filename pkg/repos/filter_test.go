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
				{Owner: "gleich", Name: "dots"},
			},
			expectedResult: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
			},
		},
		{
			repos: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
				{Owner: "cli", Name: "cli"},
			},
			expectedResult: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
			},
		},
		{
			repos: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
				{Owner: "gleich", Name: "nuke"},
			},
			expectedResult: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
			},
		},
		{
			repos: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
				{Owner: "gleich", Name: "dots"},
			},
			expectedResult: []LocalRepo{
				{Owner: "gleich", Name: "dots"},
				{Owner: "gleich", Name: "dots"},
			},
		},
	}

	for _, test := range tt {
		username := "gleich"

		result, err := FilterRepos(username, test.repos, []string{"gleich/dots"})
		assert.Equal(t, test.expectedResult, result)
		assert.NoError(t, err.Error)

		result, err = FilterRepos(username, test.repos, []string{"dots"})
		assert.Equal(t, test.expectedResult, result)
		assert.NoError(t, err.Error)

		result, err = FilterRepos(username, test.repos, []string{"gleich/dots", ""})
		assert.Equal(t, test.expectedResult, result)
		assert.NoError(t, err.Error)

	}
}
