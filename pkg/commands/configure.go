package commands

import (
	"github.com/gleich/fgh/pkg/commands/configure"
	"github.com/gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "configure",
	Short:                 "Configure fgh with an interactive prompt",
	Long:                  longDocStart + "https://github.com/gleich/fgh#%EF%B8%8F-fgh-configure",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := configure.AskQuestions()
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		err = configure.WriteConfig(config)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		statuser.Success("Wrote to config")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
