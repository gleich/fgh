package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/visualize"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var visualizeCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "visualize",
	Short:                 "Visualize all of the cloned repos in a table",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-visualize",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			config      = configuration.GetConfig(false)
			clonedRepos = reposBasedOffCustomPath(cmd, config)
		)

		ownerNameFlag, err := utils.GetBool("ownerName", cmd)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
		if ownerNameFlag {
			visualize.OutputOwnerNameList(clonedRepos)
			return
		}

		var (
			mappedRepos  = visualize.GetRepos(clonedRepos)
			createdTable = visualize.GenerateTable(mappedRepos, config)
		)
		fmt.Println(createdTable.Render())
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
	visualizeCmd.Flags().Bool("ownerName", false, "Output owner/name and path for all cloned repos")
	addCustomPathFlag(visualizeCmd)
}
