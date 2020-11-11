package migrate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawEnsureFolderExists(t *testing.T) {
	var (
		baseFolder = filepath.Join(".", "testing")
		folder     = filepath.Join(baseFolder, "testing")
	)

	err := os.MkdirAll(folder, 0777)
	assert.NoError(t, err)

	result, resultErr := rawEnsureFolderExists([]string{folder})

	err = os.RemoveAll(baseFolder)
	assert.NoError(t, err)

	assert.Equal(t, folder, result)
	assert.NoError(t, resultErr)
}
