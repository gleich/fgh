package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/Matt-Gleich/fgh/pkg/commands/login"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "login",
	Short:                 "Login to GitHub via OAuth",
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		err := login.OpenAuthPage()
		if err != nil {
			fmt.Printf("Please open the following page in your browser: %s", login.AuthPageURL())
		}
		token := login.GetToken("9000")
		configure.WriteSecrets(configure.SecretsOutline{
			PAT:      token,
			Username: login.Username(token),
		})
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
