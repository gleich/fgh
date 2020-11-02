package migrate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawEnsureFolderExists(t *testing.T) {
	folderPath := filepath.Join(".", "testing", "testing")
	err := os.MkdirAll(folderPath, 0777)
	assert.NoError(t, err)

	result, resultErr := rawEnsureFolderExists([]string{folderPath})

	err = os.RemoveAll(folderPath)
	assert.NoError(t, err)

	assert.Equal(t, folderPath, result)
	assert.NoError(t, resultErr)
}
