package commands

import (
	"fmt"

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
		cmd.Println("test")
		fmt.Println(login.GetToken("3000"))
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
