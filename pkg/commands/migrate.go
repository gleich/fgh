package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/migrate"
	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/spf13/cobra"
)

var mirgrateCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "migrate <OLD PROJECT ROOT>",
	Short:                 "Migrate your existing repos or an old fgh structure",
	Args:                  cobra.ExactArgs(1),
	// TODO: CHANGE LONG DOC
	Long: longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-ls",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			folder   = migrate.EnsureFolderExists(args)
			oldRepos = migrate.Repos(folder)
			newPaths = migrate.NewPaths(oldRepos)
		)
		migrate.ConfirmMove(newPaths)
		utils.MoveRepos(newPaths)
	},
}

func init() {
	rootCmd.AddCommand(mirgrateCmd)
}
