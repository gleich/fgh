package commands

import (
	"fmt"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fgh",
	Short: fmt.Sprintf("%v Manage your GitHub project locally", emoji.FileFolder),
	Long: fmt.Sprintf(`
%v Automate your local GitHub workspace

%v Repository: https://github.com/Matt-Gleich/fgh
%v Authors:
	- Matthew Gleich (@Matt-Gleich)

________       ______
___  __/______ ___  /_
__  /_ __  __  /_  __ \
_  __/ _  /_/ /_  / / /
/_/    _\__, / /_/ /_/
       /____/`, emoji.FileFolder, emoji.Octopus, emoji.Pager),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
