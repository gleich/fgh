package utils

import (
	"fmt"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

// Get an flag with type int
func GetInt(name string, cmd *cobra.Command) int {
	v, err := cmd.Flags().GetInt(name)
	checkErr(err, name)
	return v
}

// Get a flag with type bool
func GetBool(name string, cmd *cobra.Command) bool {
	v, err := cmd.Flags().GetBool(name)
	checkErr(err, name)
	return v
}

// Check for an error when getting a flag
func checkErr(err error, name string) {
	if err != nil {
		statuser.Error(
			fmt.Sprintf("Failed to get %v flag", name),
			err, 1,
		)
	}
}
