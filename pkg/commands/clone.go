package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone <OWNER/REPO>",
	Short:                 "Clone a repository",
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			secrets = configuration.GetSecrets()
			config  = configuration.GetConfig()
			repo    = clone.GetRepository(secrets, args)
			path    = location.RepoLocation(repo)
		)
		clone.Clone(config, secrets, repo, path)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
