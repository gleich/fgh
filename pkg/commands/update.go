package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/location"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all repos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(location.Repos())
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
