package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/visualize"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/spf13/cobra"
)

var visualizeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "visualize",
	Short:                 "Visualize the cloned repositories",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-visualize",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			config       = configuration.GetConfig(false)
			clonedRepos  = reposBasedOffCustomPath(cmd, config)
			mappedRepos  = visualize.GetRepos(clonedRepos)
			createdTable = visualize.GenerateTable(mappedRepos, config)
		)
		fmt.Println(createdTable.Render())
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
	addCustomPathFlag(visualizeCmd)
}
