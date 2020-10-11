package clean

import (
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

type Flags struct {
	Years  int
	Months int
	Days   int
}

// Parse arguments
func ParseFlags(cmd *cobra.Command) Flags {
	y, err := cmd.Flags().GetInt("years")
	if err != nil {
		statuser.Error("Failed to get years flag", err, 1)
	}

	m, err := cmd.Flags().GetInt("months")
	if err != nil {
		statuser.Error("Failed to get months flag", err, 1)
	}

	d, err := cmd.Flags().GetInt("days")
	if err != nil {
		statuser.Error("Failed to get days flag", err, 1)
	}

	return Flags{
		Years:  y,
		Months: m,
		Days:   d,
	}
}
