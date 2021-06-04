package commands

import (
	"github.com/gleich/fgh/pkg/commands/configure"
	"github.com/gleich/fgh/pkg/commands/login"
	"github.com/gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "login",
	Short:                 "Login to GitHub via OAuth",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		err := login.OpenAuthPage()
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		token := login.GetToken("9000")
		username, err := login.Username(token)
		if err.Error != nil {
			statuser.Error(err.Context, err.Error, 1)
		}

		configure.WriteSecrets(configure.SecretsOutline{
			PAT:      token,
			Username: username,
		})
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
