package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "ls <OWNER/NAME>",
	Short:                 "Get the path for cloned repo",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-ls",
	Run: func(cmd *cobra.Command, args []string) {
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

		filtered, err := repos.FilterRepos(secrets.Username, clonedRepos, args)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		fmt.Println(filtered[0].Path)
	},
	ValidArgsFunction: reposAsValidArgs,
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Allow the user to use this command on any directory
	err := addCustomPathFlag(lsCmd)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
