package clone

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/atotto/clipboard"
)

// Clone the repo
func Clone(config configure.RegularOutline, secrets configure.SecretsOutline, repo api.Repo, path string) {
	rawClone(repo, path)
	statuser.Success(fmt.Sprintf("Cloned %v/%v to:\n\t%v\n", repo.Owner, repo.Name, path))
	if config.CloneClipboard {
		err := clipboard.WriteAll(path)
		if err != nil {
			statuser.Error("Failed to copy path to clipboard", err, 1)
		}
	}
}

// Raw function for cloning the repo
func rawClone(repo api.Repo, path string) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		statuser.Error("Failed to create folder at "+path, err, 1)
	}

	gitExecPath, err := exec.LookPath("git")
	if err != nil {
		statuser.Error("Failed to locate the git executable. Please install it or put in your PATH", err, 1)
	}

	err = os.Chdir(filepath.Dir(path))
	if err != nil {
		statuser.Error("Failed to change directory to the parent folder for "+path, err, 1)
	}

	cmd := &exec.Cmd{
		Path: gitExecPath,
		Args: []string{
			gitExecPath,
			"clone",
			fmt.Sprintf("https://github.com/%v/%v.git", repo.Owner, repo.Name),
		},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	err = cmd.Run()

	if err != nil {
		statuser.Error("Failed to clone repo", err, 1)
	}
}
