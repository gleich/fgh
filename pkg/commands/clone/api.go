package clone

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
)

// Get the meta data about the repo
func GetRepository(secrets configure.SecretsOutline, args []string) (api.Repo, utils.CtxErr) {
	owner, name, err := OwnerAndName(secrets.Username, args)
	if err.Error != nil {
		return api.Repo{}, err
	}

	spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
	spin.Suffix = fmt.Sprintf(" Getting metadata for %v/%v", owner, name)
	spin.Start()

	client := api.GenerateClient(secrets.PAT)
	repo, err := api.RepoData(client, owner, name)
	if err.Error != nil {
		fmt.Println()
		return api.Repo{}, err
	}

	spin.Stop()
	statuser.Success(fmt.Sprintf("Got metadata for %v/%v\n", owner, name))
	return repo, utils.CtxErr{}
}

// Get the name of the repo and the of the owner
func OwnerAndName(username string, args []string) (string, string, utils.CtxErr) {
	var (
		owner string
		name  string
	)

	if strings.Contains(args[0], "/") {
		parts := strings.Split(args[0], "/")
		owner = parts[0]
		name = parts[1]
	} else {
		owner = username
		name = args[0]
	}

	if owner == "" {
		msg := "No owner provided"
		return "", "", utils.CtxErr{
			Context: msg,
			Error:   errors.New(msg),
		}
	}
	if name == "" {
		msg := "No repository name provided"
		return "", "", utils.CtxErr{
			Context: msg,
			Error:   errors.New(msg),
		}
	}

	return owner, name, utils.CtxErr{}
}
