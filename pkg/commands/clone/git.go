package clone

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/gleich/fgh/pkg/api"
	"github.com/gleich/fgh/pkg/commands/configure"
	"github.com/gleich/fgh/pkg/utils"
	"github.com/gleich/statuser/v2"
)

// Clone the repo
func Clone(
	config configure.RegularOutline,
	secrets configure.SecretsOutline,
	repo api.Repo,
	path string,
) utils.CtxErr {
	err := utils.CtxErr{}

	if config.SSH {
		err = sshClone(repo, path)
	} else {
		err = rawClone(repo, path)
	}

	if err.Error != nil {
		return err
	}

	statuser.Success(fmt.Sprintf("Cloned %v/%v to:\n\t%v\n", repo.Owner, repo.Name, path))
	if config.CloneClipboard {
		err := clipboard.WriteAll(path)
		if err != nil {
			return utils.CtxErr{
				Context: "Failed to copy path to clipboard",
				Error:   err,
			}
		}
	}
	return utils.CtxErr{}
}

// Raw function for cloning the repo
func rawClone(repo api.Repo, path string) utils.CtxErr {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to create folder at " + path,
			Error:   err,
		}
	}

	gitExecPath, err := exec.LookPath("git")
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to locate the git executable. Please install it or put in your PATH",
			Error:   err,
		}
	}

	err = os.Chdir(filepath.Dir(path))
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to change directory to the parent folder for " + path,
			Error:   err,
		}
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
		return utils.CtxErr{
			Context: "Failed to clone repo",
			Error:   err,
		}
	}

	return utils.CtxErr{}
}

// Raw function for cloning the repo
func sshClone(repo api.Repo, path string) utils.CtxErr {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to create folder at " + path,
			Error:   err,
		}
	}

	gitExecPath, err := exec.LookPath("git")
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to locate the git executable. Please install it or put in your PATH",
			Error:   err,
		}
	}

	err = os.Chdir(filepath.Dir(path))
	if err != nil {
		return utils.CtxErr{
			Context: "Failed to change directory to the parent folder for " + path,
			Error:   err,
		}
	}

	cmd := &exec.Cmd{
		Path: gitExecPath,
		Args: []string{
			gitExecPath,
			"clone",
			fmt.Sprintf("git@github.com:%v/%v.git", repo.Owner, repo.Name),
		},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	err = cmd.Run()

	if err != nil {
		return utils.CtxErr{
			Context: "Failed to clone repo",
			Error:   err,
		}
	}

	return utils.CtxErr{}
}
