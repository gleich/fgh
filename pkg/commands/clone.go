package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone <OWNER/REPO>",
	Short:                 "Clone a repository from GitHub",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#%EF%B8%8F-fgh-clone",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			secrets = configuration.GetSecrets()
			config  = configuration.GetConfig(false)
			repo    = clone.GetRepository(secrets, args)
			path    = repos.RepoLocation(repo, config)
		)
		clone.Clone(config, secrets, repo, path)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
