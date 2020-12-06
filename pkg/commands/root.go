package commands

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/utils"
	"github.com/Matt-Gleich/release"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fgh",
	Short: "Automate the organization of your cloned GitHub repositories",
	Long: `
   ___       __
 /'___\     /\ \
/\ \__/   __\ \ \___
\ \ ,__\/ _  \ \  _  \
 \ \ \_/\ \L\ \ \ \ \ \
  \ \_\\ \____ \ \_\ \_\
   \/_/ \/___L\ \/_/\/_/
          /\____/
	  \_/__/

Automate the organization of your cloned GitHub repositories

Repository: https://github.com/Matt-Gleich/fgh
Authors:
  - Matthew Gleich (@Matt-Gleich)
  - Caleb Denio (@cjdenio)
  - Safin Singh (@safinsingh)`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag := utils.GetBool("version", cmd)
		if versionFlag {
			version := "v2.3.1"

			spin := spinner.New(utils.SpinnerCharSet, utils.SpinnerSpeed)
			spin.Suffix = " Checking for update"
			spin.Start()
			outdated, v, err := release.Check(version, "https://github.com/Matt-Gleich/fgh")
			spin.Stop()

			if err != nil {
				statuser.Error("Failed to get latest version of fgh", err, 1)
			}
			if outdated {
				statuser.Warning(fmt.Sprintf("%v of fgh is out! Please upgrade.", v))
			} else {
				fmt.Println("You are on the latest version of fgh.")
			}
			fmt.Println(version)
		} else {
			err := cmd.Help()
			if err != nil {
				statuser.Error("Failed to display help", err, 1)
			}
			os.Exit(0)
		}
	},
}

func Execute() {
	rootCmd.Flags().Bool("version", false, "Get the current version of fgh and check for an update")
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
