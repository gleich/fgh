package clean

import (
	"github.com/Matt-Gleich/fgh/pkg/utils"
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
		Years:        utils.GetInt("years", cmd),
		Months:       utils.GetInt("months", cmd),
		Days:         utils.GetInt("days", cmd),
		SkipOutdated: utils.GetBool("skipOutdated", cmd),
		SkipDeleted:  utils.GetBool("skipDeleted", cmd),
	}
}
