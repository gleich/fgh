package clone

import (
	"fmt"
	"strings"
	"time"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"github.com/enescakir/emoji"
)

// Get the meta data about the repo
func GetRepository(secrets configure.SecretsOutline, args []string) api.Repo {
	owner, name := OwnerAndName(secrets, args)
	spin := spinner.New(spinner.CharSets[1], 40*time.Millisecond)
	spin.Suffix = fmt.Sprintf(" %v  Getting metadata for %v/%v", emoji.Information, owner, name)
	spin.Start()

	client := api.GenerateClient()
	repo, err := api.RepoData(client, owner, name)
	if err != nil {
		statuser.Error("Failed to get repo information", err, 1)
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("Got metadata for %v/%v\n", owner, name))
	return repo
}

// Get the name of the repo and the of the owner
func OwnerAndName(secrets configure.SecretsOutline, args []string) (owner string, name string) {
	if strings.Contains(args[0], "/") {
		parts := strings.Split(args[0], "/")
		owner = parts[0]
		name = parts[1]
	} else {
		owner = secrets.Username
		name = args[0]
	}

	if owner == "" {
		statuser.ErrorMsg("No owner provided", 1)
	}
	if name == "" {
		statuser.ErrorMsg("No repository name provided", 1)
	}

	return owner, name
}
