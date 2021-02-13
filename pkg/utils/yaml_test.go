package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteReadYAML(t *testing.T) {
	var (
		data  = map[string]string{"test": "test"}
		fName = "test.yaml"
	)

	// Writing to the file
	assert.Nil(t, WriteYAML(data, fName).Error)

	// Reading from the file
	var readData struct {
		Test string `yaml:"test"`
	}
	readErr := ReadYAML(fName, &readData)

	// Removing the created file
	assert.Nil(t, os.Remove(fName))

	assert.Nil(t, readErr.Error)
	assert.Equal(t, "test", readData.Test)
}
