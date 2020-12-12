package commands

import (
	"fmt"
	"strings"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
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
func addCustomPathFlag(cmd *cobra.Command) {
	var (
		config     = configuration.GetConfig(true)
		rootFolder = repos.StructureRootPath(config)
	)
	cmd.Flags().StringP(
		pathFlagName,
		"p",
		"", // This would be set as the rootFolder but we need to know if the user actually gave the path and therefore needs to do a brute force search with repos.Repos().
		fmt.Sprintf("Root folder where the repo(s) should be in at some level. Relative to where you are running this command (default \"%v\")", rootFolder),
	)
}

// Get the value for the custom path
func reposBasedOffCustomPath(cmd *cobra.Command, config configure.RegularOutline) []repos.LocalRepo {
	path, err := cmd.Flags().GetString(pathFlagName)
	if err != nil {
		statuser.Error("Failed to get custom path flag", err, 1)
	}

	if path == "" {
		return repos.ReposInStructure(config, true)
	}
	return repos.Repos(path, true)
}

// Set the valid args as the local repos.
func reposAsValidArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var (
		secrets     = configuration.GetSecrets()
		config      = configuration.GetConfig(false)
		clonedRepos = reposBasedOffCustomPath(cmd, config)
		repoPairs   = []string{}
	)

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
