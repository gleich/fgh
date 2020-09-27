package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone <OWNER/REPO>",
	Short:                 "Clone a repository",
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		secrets := configuration.GetSecrets()
		repo := clone.GetRepository(secrets, args)
		path := clone.Location(repo)
		fmt.Println(path)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
