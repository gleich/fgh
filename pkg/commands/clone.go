package commands

import (
	"github.com/Matt-Gleich/fgh/pkg/commands/clone"
	"github.com/Matt-Gleich/fgh/pkg/configuration"
	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "clone <OWNER/REPO>",
	Short:                 "Clone a repository from GitHub.",
	Args:                  cobra.ExactArgs(1),
	Long: `It all starts with cloning a repo which is done when you run fgh clone OWNER/NAME or fgh clone NAME if the repo is under your account. All repos are cloned in the following structure:

~
└──github
   └── OWNER
       └── TYPE
           └── MAIN LANGUAGE
               └── NAME


- OWNER: The owner of the repo.
- TYPE: The type of the repo; one of the following:
  - public
  - private
  - template
  - archived
  - disabled
  - mirror
  - fork
- MAIN LANGUAGE: The main language for the repo. If no language is detected then will just be Other.
- Name: The name of the repo.

So if you were to clone this repo using fgh clone Matt-Gleich/fgh it would be cloned to the to ~/github/Matt-Gleich/public/Go/fgh/.

Once you are done cloning the repo you can have the path copied to your clipboard automatically. If you are on Linux you need xclip or xsel for this to work.

This structure can be somewhat difficult to navigate in the terminal using conventional methods. I suggest navigators such as ranger (https://github.com/ranger/ranger) to help speed up the process.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			secrets = configuration.GetSecrets()
			config  = configuration.GetConfig()
			repo    = clone.GetRepository(secrets, args)
			path    = location.RepoLocation(repo)
		)
		clone.Clone(config, secrets, repo, path)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
