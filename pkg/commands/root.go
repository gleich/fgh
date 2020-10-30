package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/release"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

// Starter for all long form docs
const longDocStart = "\nDocumentation for this subcommand: "

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fgh",
	Short: fmt.Sprintf("%v Automate the organization of your cloned GitHub repositories", emoji.FileFolder),
	Long: fmt.Sprintf(`
   ___       __
 /'___\     /\ \
/\ \__/   __\ \ \___
\ \ ,__\/ _  \ \  _  \
 \ \ \_/\ \L\ \ \ \ \ \
  \ \_\\ \____ \ \_\ \_\
   \/_/ \/___L\ \/_/\/_/
          /\____/
	  \_/__/

%v Automate the organization of your cloned GitHub repositories

%v Repository: https://github.com/Matt-Gleich/fgh
%v Authors:
	- Matthew Gleich (@Matt-Gleich)
	- Caleb Denio (@cjdenio)
	- Safin Singh (@safinsingh)`, emoji.FileFolder, emoji.Octopus, emoji.Pager),
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag := utils.GetBool("version", cmd)
		if versionFlag {
			version := "v1.4.0"
			outdated, v, err := release.Check(version, "https://github.com/Matt-Gleich/fgh")
			if err != nil {
				statuser.Error("Failed to get latest version of fgh", err, 1)
			}
			if outdated {
				statuser.Warning(fmt.Sprintf("%v of fgh is out! Please upgrade.", v))
			}
			fmt.Println(version)
		}
	},
}

func Execute() {
	rootCmd.Flags().Bool("version", false, "Get the current version of fgh and check for an update")
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
