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
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		filtered := repos.FilterRepos(configuration.GetSecrets().Username, repos.Repos(), args)
		fmt.Println(filtered[0].Path)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
