package commands

import (
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion for fgh",
	Long: `# You will need to start a new shell for this setup to take effect.

Fish:

$ yourprogram completion fish | source

# To load completions for each session, execute once:
$ yourprogram completion fish > ~/.config/fish/completions/yourprogram.fish
`,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
		if err != nil {
			statuser.Error("Failed to generate completion for "+args[0], err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
