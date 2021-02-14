package repos

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGitRepo(t *testing.T) {
	tests := []struct {
		extraFolders []string
	}{
		{extraFolders: []string{"folder-1/.git", "folder-1/other_folder"}},
		{extraFolders: []string{"folder-1/.git", "folder-2/.git"}},
		{extraFolders: []string{"folder-1/child-folder", "folder-1/.git", "folder-2"}},
	}

	for _, tt := range tests {
		// Creating folders & files
		for _, folder := range tt.extraFolders {
			err := os.MkdirAll(filepath.FromSlash(folder), 0777)
			assert.NoError(t, err)
		}

		isRepo, err := IsGitRepo("folder-1")
		assert.NoError(t, err.Error)
		assert.True(t, isRepo)

		// Removing files and folders
		for _, folder := range tt.extraFolders {
			for _, childFolder := range strings.Split(folder, string(filepath.Separator)) {
				err := os.RemoveAll(childFolder)
				assert.NoError(t, err)
			}
		}
	}
}
