package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Get an flag with type int
func GetInt(name string, cmd *cobra.Command) (int, CtxErr) {
	v, err := cmd.Flags().GetInt(name)
	return v, CtxErr{
		Error:   err,
		Context: fmt.Sprintf("Failed to get %v flag (type int)", name),
	}
}

// Get a flag with type bool
func GetBool(name string, cmd *cobra.Command) (bool, CtxErr) {
	v, err := cmd.Flags().GetBool(name)
	return v, CtxErr{
		Error:   err,
		Context: fmt.Sprintf("Failed to get %v flag (type bool)", name),
	}
}
