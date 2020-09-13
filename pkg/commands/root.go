package commands

import (
	"github.com/Matt-Gleich/statuser"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fgh",
	Short: "📁 Manage your GitHub project locally",
	Long: `
📁 Manage your GitHub project locally

🐙 Repository: https://github.com/Matt-Gleich/fgh
📟 Authors:
	- Matthew Gleich (@Matt-Gleich)

________       ______
___  __/______ ___  /_
__  /_ __  __  /_  __ \
_  __/ _  /_/ /_  / / /
/_/    _\__, / /_/ /_/
       /____/`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
