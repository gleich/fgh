package clean

import (
	"fmt"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

type Flags struct {
	Years        int
	Months       int
	Days         int
	SkipOutdated bool
	SkipDeleted  bool
}

// Parse arguments
func ParseFlags(cmd *cobra.Command) Flags {
	return Flags{
		Years:        getInt("years", cmd),
		Months:       getInt("months", cmd),
		Days:         getInt("days", cmd),
		SkipOutdated: getBool("skipOutdated", cmd),
		SkipDeleted:  getBool("skipDeleted", cmd),
	}
}

// Get an flag with type int
func getInt(name string, cmd *cobra.Command) int {
	v, err := cmd.Flags().GetInt(name)
	checkErr(err, name)
	return v
}

// Get a flag with type bool
func getBool(name string, cmd *cobra.Command) bool {
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
