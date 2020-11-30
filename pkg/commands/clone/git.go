package clone

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Clone the repo
func Clone(config configure.RegularOutline, secrets configure.SecretsOutline, repo api.Repo, path string) {
	rawClone(secrets, repo, path)
	statuser.Success(fmt.Sprintf("Cloned %v/%v to:\n\t%v\n", repo.Owner, repo.Name, path))
	if config.CloneClipboard {
		err := clipboard.WriteAll(path)
		if err != nil {
			statuser.Error("Failed to copy path to clipboard", err, 1)
		}
	}
}

// Raw function for cloning the repo
func rawClone(secrets configure.SecretsOutline, repo api.Repo, path string) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		statuser.Error("Failed to create folder at "+path, err, 1)
	}

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(" Cloning %v/%v", repo.Owner, repo.Name)
	spin.Start()

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: fmt.Sprintf("https://github.com/%v/%v.git", repo.Owner, repo.Name),
		Auth: &http.BasicAuth{
			Username: secrets.Username,
			Password: secrets.PAT,
		},
	})

	spin.Stop()
	if err != nil {
		statuser.Error("Failed to clone repo", err, 1)
	}
}
