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
	err := WriteYAML(data, fName)
	assert.Nil(t, err)

	// Reading from the file
	var readData struct {
		Test string `yaml:"test"`
	}
	readErr := ReadYAML(fName, &readData)

	// Removing the created file
	err = os.Remove(fName)
	assert.Nil(t, err)

	assert.Nil(t, readErr)
	assert.Equal(t, "test", readData.Test)
}
