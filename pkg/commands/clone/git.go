package clone

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/atotto/clipboard"
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

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: fmt.Sprintf("https://github.com/%v/%v.git", repo.Owner, repo.Name),
		Auth: &http.BasicAuth{
			Username: api.Username(),
			Password: secrets.PAT,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		statuser.Error("Failed to clone repo", err, 1)
	}
}
