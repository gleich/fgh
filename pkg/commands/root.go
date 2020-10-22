package commands

import (
	"fmt"

	"github.com/Matt-Gleich/release"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fgh",
	Short: fmt.Sprintf("%v Manage your GitHub project locally", emoji.FileFolder),
	Long: fmt.Sprintf(`
%v Automate your the organization of your cloned GitHub repos

%v Repository: https://github.com/Matt-Gleich/fgh
%v Authors:
	- Matthew Gleich (@Matt-Gleich)

   ___       __
 /'___\     /\ \
/\ \__/   __\ \ \___
\ \ ,__\/ _  \ \  _  \
 \ \ \_/\ \L\ \ \ \ \ \
  \ \_\\ \____ \ \_\ \_\
   \/_/ \/___L\ \/_/\/_/
          /\____/
          \_/__/`, emoji.FileFolder, emoji.Octopus, emoji.Pager),
	Run: func(cmd *cobra.Command, args []string) {
		isOutdated, version, _ := release.Check("v1.0.0", "https://github.com/Matt-Gleich/fgh")
		if isOutdated {
			statuser.Warning(fmt.Sprintf(
				"Version %v of fgh is now available! Please update at your convenience.",
				version,
			))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
