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
			username:      "Matt-Gleich",
			args:          []string{"dots"},
			expectedOwner: "Matt-Gleich",
			expectedName:  "dots",
		},
		{
			username:      "Matt-Gleich",
			args:          []string{"Matt-Gleich/dots"},
			expectedOwner: "Matt-Gleich",
			expectedName:  "dots",
		},
		{
			username:      "nat",
			args:          []string{"Matt-Gleich/dots"},
			expectedOwner: "Matt-Gleich",
			expectedName:  "dots",
		},
	}

	for _, test := range tt {
		ownerResult, nameResult := OwnerAndName(test.username, test.args)
		assert.Equal(t, test.expectedOwner, ownerResult)
		assert.Equal(t, test.expectedName, nameResult)
	}
}
