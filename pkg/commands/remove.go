package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/remove"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
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
		var (
			secrets     = configuration.GetSecrets()
			config      = configuration.GetConfig()
			clonedRepos = repos.Repos(config)
		)
		filtered := repos.FilterRepos(secrets.Username, clonedRepos, args)
		remove.RemoveRepos(filtered)
		clean.CleanUp(config)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
