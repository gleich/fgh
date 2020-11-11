package commands

import (
	"fmt"

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
		config     = configuration.GetConfig().StructureRoot
		rootFolder = repos.StructureRootFolder(config)
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
		return repos.ReposInStructure(config)
	}
	return repos.Repos(path)
}
