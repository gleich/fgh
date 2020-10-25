package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/api"
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/remove"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "remove <OWNER/NAME>",
	Short:                 fmt.Sprintf("%v  Remove a cloned repo", emoji.Wastebasket),
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		filtered := remove.FilterRepos(api.Username(), repos.Repos(), args)
		remove.RemoveRepos(filtered)
		clean.CleanUp()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
