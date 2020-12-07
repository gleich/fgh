package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/tree"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "tree",
	Short:                 "Visualize the cloned repositories",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-tree",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			config      = configuration.GetConfig(false)
			clonedRepos = reposBasedOffCustomPath(cmd, config)
		)
		fmt.Printf("%#v\n", tree.GetRepos(clonedRepos))
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
	addCustomPathFlag(treeCmd)
}
