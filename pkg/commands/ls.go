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
		var (
			secrets     = configuration.GetSecrets()
			config      = configuration.GetConfig(false)
			clonedRepos = reposBasedOffCustomPath(cmd, config)
		)

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
	addCustomPathFlag(lsCmd)
}
