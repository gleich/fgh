package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clean"
	"github.com/Matt-Gleich/fgh/pkg/commands/remove"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a cloned repo.",
	Run: func(cmd *cobra.Command, args []string) {
		filtered := remove.FilterRepos(configuration.GetSecrets(), location.Repos(), args)
		remove.RemoveRepos(filtered)
		clean.CleanUp()
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
