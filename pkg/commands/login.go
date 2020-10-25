package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"

	"github.com/Matt-Gleich/fgh/pkg/commands/login"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "login",
	Short:                 fmt.Sprintf("%v  Log in to GitHub via OAuth", emoji.Locked),
	Args:                  cobra.NoArgs,
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#-fgh-remove",
	Run: func(cmd *cobra.Command, args []string) {
		err := login.OpenAuthPage()
		if err != nil {
			fmt.Printf("Please open the following page in your browser: %s", login.AuthPageURL())
		}
		token := login.GetToken("9000")
		configure.WriteSecrets(configure.CreateFolders(), configure.SecretsOutline{
			PAT: token,
		})
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
