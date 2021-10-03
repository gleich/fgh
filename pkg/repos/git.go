package repos

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/fgh/pkg/utils"
	"github.com/go-git/go-git/v5"
)

// Get the owner and name of the repo just from the default remote
func OwnerAndNameFromRemote(path string) (string, string, utils.CtxErr) {
	var (
		owner string
		name  string
	)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", "", utils.CtxErr{
			Context: "Failed to read from git repo located in " + path,
			Error:   err,
		}
	}

	remotes, err := repo.Remotes()
	if err != nil {
		return "", "", utils.CtxErr{
			Context: "Failed to get remotes from git repo located in " + path,
			Error:   err,
		}
	}

	// If origin isn't a remote then ask the user to see which one is
	var (
		foundOrigin = false
		url         string
	)
	for _, remote := range remotes {
		if remote.Config().Name == "origin" {
			foundOrigin = true
			defaultURL, err := getDefaultURL(remote, path)
			url = defaultURL
			if err.Error != nil {
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
			return "", "", utils.CtxErr{
				Context: "Failed to ask what remote is the default remote",
				Error:   err,
			}
		}

		defaultURL, errCtx := getDefaultURL(remoteOptionMap[defaultRemoteStr], path)
		url = defaultURL
		if err != nil {
			return "", "", errCtx
		}
	}

	// Getting name and owner
	parts := strings.Split(url, "/")
	owner = strings.TrimPrefix(parts[len(parts)-2], "git@github.com:")
	name = strings.TrimSuffix(parts[len(parts)-1], ".git")

	return owner, name, utils.CtxErr{}
}

// If the number of urls for a remote is over 1 then ask the user
// which one to use. If not, just return the first url.
func getDefaultURL(remote *git.Remote, path string) (string, utils.CtxErr) {
	urls := remote.Config().URLs
	if len(urls) == 0 {
		errMsg := "No remotes found for " + path
		return "", utils.CtxErr{
			Context: errMsg,
			Error:   errors.New(errMsg),
		}
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
			return "", utils.CtxErr{
				Context: "Failed to ask what the default remote url is",
				Error:   err,
			}
		}
		return chosenURL, utils.CtxErr{}
	}
	return urls[0], utils.CtxErr{}
}

// Check so see if a repo has a dirty working tree or any commits that haven't been pushed
// Returns if all changes have been committed and pushed
func WorkingState(path string) (bool, bool, utils.CtxErr) {
	var (
		committed bool
		pushed    bool
	)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return false, false, utils.CtxErr{
			Context: "Failed to read from git repo located in " + path,
			Error:   err,
		}
	}

	// Working tree staged and unstaged changes
	workingTree, err := repo.Worktree()
	if err != nil {
		return false, false, utils.CtxErr{
			Context: "Failed to get working tree for " + path,
			Error:   err,
		}
	}

	status, err := workingTree.Status()
	if err != nil {
		return false, false, utils.CtxErr{
			Context: "Failed to get status of changes for " + path,
			Error:   err,
		}
	}
	if len(status) == 0 {
		committed = true
	}

	// Commits not pushed
	err = os.Chdir(path)
	if err != nil {
		return false, false, utils.CtxErr{
			Context: "Failed to change directory into " + path,
			Error:   err,
		}
	}

	gitPath, err := exec.LookPath("git")
	if err != nil {
		return false, false, utils.CtxErr{
			Context: "Looks like you don't have git installed. Please install it.",
			Error:   err,
		}
	}

	pushed, errCtx := runGitCherry(gitPath, path)
	if errCtx.Error != nil {
		// Refreshing origin
		err = os.RemoveAll("./.git/refs/remotes/origin")
		if err != nil {
			return false, false, utils.CtxErr{
				Context: "Failed to remove local reference of origin remote",
				Error:   err,
			}
		}

		err = exec.Command(gitPath, "fetch", "--all").Run()
		if err != nil {
			return false, false, utils.CtxErr{
				Context: "Failed to fetch latest commits" + path,
				Error:   err,
			}
		}

		pushed, errCtx = runGitCherry(gitPath, path)
		if errCtx.Error != nil {
			return false, false, errCtx
		}
	}

	return committed, pushed, utils.CtxErr{}
}

// Run git cherry to see if there are any commits not yet pushed
func runGitCherry(gitPath string, path string) (bool, utils.CtxErr) {
	out, err := exec.Command(gitPath, "cherry", "-v").Output()
	if err != nil {
		return false, utils.CtxErr{
			Context: "Failed to check if repo has any commits not pushed. Location: " + path,
			Error:   err,
		}
	}

	if len((strings.Split(string(out), "\n"))) != 1 {
		return false, utils.CtxErr{}
	}
	return true, utils.CtxErr{}
}

// Checks to make sure the given folder has a .git folder inside
func IsGitRepo(path string) (bool, utils.CtxErr) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, utils.CtxErr{
			Context: "Failed to list " + path,
			Error:   err,
		}
	}

	for _, f := range files {
		if f.IsDir() && f.Name() == ".git" {
			return true, utils.CtxErr{}
		}
	}
	return false, utils.CtxErr{}
}
