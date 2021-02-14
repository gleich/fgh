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
		config, err := configuration.GetConfig(false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		clonedRepos, err := reposBasedOffCustomPath(cmd, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		ownerNameFlag, err := utils.GetBool("ownerName", cmd)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
		if ownerNameFlag {
			visualize.OutputOwnerNameList(clonedRepos)
			return
		}

		mappedRepos, err := visualize.GetRepos(clonedRepos)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		createdTable := visualize.GenerateTable(mappedRepos, config)
		fmt.Println(createdTable.Render())
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
	visualizeCmd.Flags().Bool("ownerName", false, "Output owner/name and path for all cloned repos")

	// Allow the user to use this command on any directory
	err := addCustomPathFlag(visualizeCmd)
	if err.Error != nil {
		statuser.Error(err.Context, err.Error, 1)
	}
}
