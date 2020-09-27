package clone

import (
	"fmt"
	"os"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func Clone(secrets configure.SecretsOutline, repo Repository, path string) {
	spin := spinner.New(spinner.CharSets[13], 40*time.Millisecond)
	spin.Suffix = fmt.Sprintf(" ☁️  Cloning %v/%v", repo.Owner, repo.Name)
	spin.Start()
	rawClone(secrets, repo, path)
	spin.Stop()
	statuser.Success(fmt.Sprintf("Cloned %v/%v to:\n\t%v", repo.Owner, repo.Name, path))
}

// Raw function for cloning the repo
func rawClone(secrets configure.SecretsOutline, repo Repository, path string) {
	// Creating folder location:
	err := os.MkdirAll(path, 0777)
	if err != nil {
		statuser.Error("Failed to create folder at "+path, err, 1)
	}

	_, err = git.PlainClone(path, false, &git.CloneOptions{
		URL: fmt.Sprintf("https://github.com/%v/%v.git", repo.Owner, repo.Name),
		Auth: &http.BasicAuth{
			Username: secrets.Username,
			Password: secrets.PAT,
		},
	})
	if err != nil {
		statuser.Error("Failed to clone repo", err, 1)
	}
}
