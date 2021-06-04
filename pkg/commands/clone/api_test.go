package clone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOwnerAndName(t *testing.T) {
	tt := []struct {
		username      string
		args          []string
		expectedOwner string
		expectedName  string
	}{
		{
			username:      "gleich",
			args:          []string{"dots"},
			expectedOwner: "gleich",
			expectedName:  "dots",
		},
		{
			username:      "gleich",
			args:          []string{"gleich/dots"},
			expectedOwner: "gleich",
			expectedName:  "dots",
		},
		{
			username:      "nat",
			args:          []string{"gleich/dots"},
			expectedOwner: "gleich",
			expectedName:  "dots",
		},
	}

	for _, test := range tt {
		ownerResult, nameResult, err := OwnerAndName(test.username, test.args)
		assert.NoError(t, err.Error)
		assert.Equal(t, test.expectedOwner, ownerResult)
		assert.Equal(t, test.expectedName, nameResult)
	}
}
