package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/repos"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone <OWNER/REPO>",
	Short:                 "Clone a repository from GitHub",
	Args:                  cobra.ExactArgs(1),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#%EF%B8%8F-fgh-clone",
	Run: func(cmd *cobra.Command, args []string) {
		secrets, err := configuration.GetSecrets()
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		config, err := configuration.GetConfig(false)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		repo, err := clone.GetRepository(secrets, args)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		path, err := repos.RepoLocation(repo, config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = clone.Clone(config, secrets, repo, path)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
