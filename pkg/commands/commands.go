package commands

import (
	"fmt"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

const (
	// Starter for all long form docs
	longDocStart = "\nDocumentation for this subcommand: "
	// Name for the custom path flag
	pathFlagName = "path"
)

// Add the custom path flag to the command
func addCustomPathFlag(cmd *cobra.Command) utils.CtxErr {
	config, err := configuration.GetConfig(true)
	if err.Error != nil {
		return err
	}
	rootFolder, err := repos.StructureRootPath(config)

	cmd.Flags().StringP(
		pathFlagName,
		"p",
		"", // This would be set as the rootFolder but we need to know if the user actually gave the path and therefore needs to do a brute force search with repos.Repos().
		fmt.Sprintf("Root folder where the repo(s) should be in at some level. Relative to where you are running this command (default \"%v\")", rootFolder),
	)
	return utils.CtxErr{}
}

// Get the value for the custom path
func reposBasedOffCustomPath(cmd *cobra.Command, config configure.RegularOutline) ([]repos.LocalRepo, utils.CtxErr) {
	path, err := cmd.Flags().GetString(pathFlagName)
	if err != nil {
		return []repos.LocalRepo{}, utils.CtxErr{
			Context: fmt.Sprintf("Failed to get %v flag", pathFlagName),
			Error:   err,
		}
	}

	if path == "" {
		clonedRepos, _ := repos.ReposInStructure(config, true)
		return clonedRepos, utils.CtxErr{
			Context: "Failed to get custom path flag",
			Error:   err,
		}
	}

	clonedRepos, _ := repos.Repos(path, true)
	return clonedRepos, utils.CtxErr{}
}

// Set the valid args as the local repos.
func reposAsValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	secrets, err := configuration.GetSecrets()
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}

	config, err := configuration.GetConfig(false)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}

	clonedRepos, err := reposBasedOffCustomPath(cmd, config)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}

	var repoPairs = []string{}
	for _, repo := range clonedRepos {
		if repo.Owner == secrets.Username {
			repoPairs = append(repoPairs, repo.Name)
			continue
		}
		pair := repo.Owner + "/" + repo.Name
		repoPairs = append(repoPairs, pair)
	}
	if toComplete != "" {
		repoPairsCleaned := []string{}
		for _, repo := range repoPairs {
			if strings.HasPrefix(repo, toComplete) {
				repoPairsCleaned = append(repoPairsCleaned, repo)
			}
		}
		return repoPairsCleaned, cobra.ShellCompDirectiveNoSpace
	}

	return repoPairs, cobra.ShellCompDirectiveDefault
}
