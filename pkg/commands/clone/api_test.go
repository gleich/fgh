package clone

import (
	"testing"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/stretchr/testify/assert"
)

func TestOwnerAndName(t *testing.T) {
	tt := []struct {
		secrets       configure.SecretsOutline
		args          []string
		expectedOwner string
		expectedName  string
	}{
		{
			secrets:       configure.SecretsOutline{Username: "Matt-Gleich"},
			args:          []string{"dots"},
			expectedOwner: "Matt-Gleich",
			expectedName:  "dots",
		},
		{
			secrets:       configure.SecretsOutline{Username: "Matt-Gleich"},
			args:          []string{"Matt-Gleich/dots"},
			expectedOwner: "Matt-Gleich",
			expectedName:  "dots",
		},
		{
			secrets:       configure.SecretsOutline{Username: "nat"},
			args:          []string{"Matt-Gleich/dots"},
			expectedOwner: "Matt-Gleich",
			expectedName:  "dots",
		},
	}

	for _, test := range tt {
		ownerResult, nameResult := ownerAndName(test.secrets, test.args)
		assert.Equal(t, test.expectedOwner, ownerResult)
		assert.Equal(t, test.expectedName, nameResult)
	}
}
