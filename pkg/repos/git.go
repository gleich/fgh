package repos

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/go-git/go-git/v5"
)

// Check so see if a repo has a dirty working tree or any commits that haven't been pushed
// Returns if all changes have been committed and pushed
func WorkingState(path string) (committed bool, pushed bool) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		statuser.Error("Failed to read from git repo located in "+path, err, 1)
	}

	// Working tree staged and unstaged changes
	workingTree, err := repo.Worktree()
	if err != nil {
		statuser.Error("Failed to get working tree for "+path, err, 1)
	}
	status, err := workingTree.Status()
	if err != nil {
		statuser.Error("Failed to get status of changes for "+path, err, 1)
	}
	if len(status) == 0 {
		committed = true
	}

	// Commits not pushed
	err = os.Chdir(path)
	if err != nil {
		statuser.Error("Failed to change directory into "+path, err, 1)
	}
	_, err = exec.LookPath("git")
	if err != nil {
		statuser.Error("Looks like you don't have git installed. Please install it.", err, 1)
	}
	out, err := exec.Command("git", "cherry", "-v").Output()
	if err != nil {
		statuser.Error("Failed to check if repo has any commits not pushed. Location: "+path, err, 1)
	}
	if len((strings.Split(string(out), "\n"))) == 1 {
		pushed = true
	}

	return committed, pushed
}

// Checks to make sure the given folder has a .git folder inside
func isGitRepo(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		cwd, err1 := os.Getwd()
		if err1 != nil {
			statuser.Error("Failed to get current working directory", err, 1)
		}
		statuser.Error("Failed to list "+cwd, err, 1)
	}

	for _, f := range files {
		if f.IsDir() && f.Name() == ".git" {
			return true
		}
	}
	return false
}
