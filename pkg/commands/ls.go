package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "ls <OWNER/NAME>",
	Short:                 fmt.Sprintf("%v Get the path for cloned repo", emoji.Compass),
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-ls",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			secrets     = configuration.GetSecrets()
			config      = configuration.GetConfig()
			clonedRepos = repos.Repos(config)
		)
		filtered := repos.FilterRepos(secrets.Username, clonedRepos, args)
		fmt.Println(filtered[0].Path)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
