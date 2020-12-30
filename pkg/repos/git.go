package repos

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/go-git/go-git/v5"
)

// Get the owner and name of the repo just from the default remote
func OwnerAndNameFromRemote(path string) (owner string, name string, err error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		statuser.Error("Failed to read from git repo located in "+path, err, 1)
	}
	remotes, err := repo.Remotes()
	if err != nil {
		statuser.Error("Failed to get remotes for git repo located in "+path, err, 1)
	}

	// If origin isn't a remote then ask the user to see which one is
	var (
		foundOrigin = false
		url         string
	)
	for _, remote := range remotes {
		if remote.Config().Name == "origin" {
			foundOrigin = true
			url, err = getDefaultURL(remote, path)
			if err != nil {
				return "", "", err
			}
		}
	}
	if !foundOrigin {
		var (
			remoteOptionMap = map[string]*git.Remote{}
			remoteOptions   = []string{}
		)
		for _, remote := range remotes {
			strRemote := fmt.Sprint(remote)
			remoteOptionMap[strRemote] = remote
			remoteOptions = append(remoteOptions, strRemote)
		}

		// Asking the user
		var defaultRemoteStr string
		prompt := &survey.Select{
			Message: fmt.Sprintf("Which remote is the default remote for %v?", path),
			Options: remoteOptions,
		}
		err := survey.AskOne(prompt, &defaultRemoteStr)
		if err != nil {
			return "", "", err
		}
		url, err = getDefaultURL(remoteOptionMap[defaultRemoteStr], path)
		if err != nil {
			return "", "", err
		}
	}

	// Getting name and owner
	parts := strings.Split(url, "/")
	owner = parts[len(parts)-2]
	name = strings.TrimSuffix(parts[len(parts)-1], ".git")

	return owner, name, nil
}

// If the number of urls for a remote is over 1 then ask the user
// which one to use. If not, just return the first url.
func getDefaultURL(remote *git.Remote, path string) (string, error) {
	urls := remote.Config().URLs
	if len(urls) == 0 {
		return "", errors.New("No remotes found for " + path)
	}
	if len(urls) > 1 {
		var chosenURL string
		prompt := &survey.Select{
			Message: fmt.Sprintf(
				"What is the default url for %v in %v?",
				remote.Config().Name,
				path,
			),
			Options: urls,
		}

		err := survey.AskOne(prompt, &chosenURL)
		if err != nil {
			return "", err
		}
		return chosenURL, nil
	}
	return urls[0], nil
}

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

	gitPath, err := exec.LookPath("git")
	if err != nil {
		statuser.Error("Looks like you don't have git installed. Please install it.", err, 1)
	}

	cmd := exec.Command(gitPath, "cherry", "-v")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		statuser.Error("Failed to check if repo has any commits not pushed. Location: "+path, err, 1)
	}
	if len((strings.Split(string(out), "\n"))) == 1 {
		pushed = true
	}

	return committed, pushed
}

// Checks to make sure the given folder has a .git folder inside
func IsGitRepo(path string) bool {
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
