package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
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
			config      = configuration.GetConfig()
			clonedRepos = reposBasedOffCustomPath(cmd, config)
		)
		filtered := repos.FilterRepos(secrets.Username, clonedRepos, args)
		fmt.Println(filtered[0].Path)
	},
	ValidArgsFunction: validArgsAsRepos,
}

func init() {
	rootCmd.AddCommand(lsCmd)
	addCustomPathFlag(lsCmd)
}
