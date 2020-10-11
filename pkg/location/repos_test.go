package location

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirs(t *testing.T) {
	tests := []struct {
		folders       []string
		files         []string
		valid_folders []string
	}{
		{
			folders:       []string{"folder-1"},
			files:         []string{"file1.txt"},
			valid_folders: []string{"folder-1"},
		},
		{
			folders:       []string{"folder-1", "folder-2"},
			files:         []string{"file1.txt"},
			valid_folders: []string{"folder-1", "folder-2"},
		},
		{
			folders:       []string{filepath.Join("folder-1", "child-folder"), "folder-2"},
			files:         []string{filepath.Join("folder-2", "file1.txt")},
			valid_folders: []string{"folder-1", "folder-2"},
		},
		{
			folders:       []string{},
			files:         []string{},
			valid_folders: []string(nil),
		},
	}

	for _, tt := range tests {
		// Creating folders & files
		for _, folder := range tt.folders {
			err := os.MkdirAll(folder, 0777)
			assert.NoError(t, err)
		}
		for _, file := range tt.files {
			err := os.MkdirAll(filepath.Dir(file), 0777)
			assert.NoError(t, err)
			_, err = os.Create(file)
			assert.NoError(t, err)
		}

		assert.Equal(t, tt.valid_folders, dirs())

		// Removeing files and folders
		for _, folder := range append(tt.folders, tt.files...) {
			for _, childFolder := range strings.Split(folder, string(filepath.Separator)) {
				err := os.RemoveAll(childFolder)
				assert.NoError(t, err)
			}
		}
	}
}
