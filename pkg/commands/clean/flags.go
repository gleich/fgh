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
func ParseFlags(cmd *cobra.Command) (Flags, utils.CtxErr) {
	years, err := utils.GetInt("years", cmd)
	if err.Error != nil {
		return Flags{}, err
	}

	months, err := utils.GetInt("months", cmd)
	if err.Error != nil {
		return Flags{}, err
	}

	days, err := utils.GetInt("days", cmd)
	if err.Error != nil {
		return Flags{}, err
	}

	skipOutdated, err := utils.GetBool("skipOutdated", cmd)
	if err.Error != nil {
		return Flags{}, err
	}

	skipDeleted, err := utils.GetBool("skipDeleted", cmd)
	if err.Error != nil {
		return Flags{}, err
	}

	return Flags{
		Years:        years,
		Months:       months,
		Days:         days,
		SkipOutdated: skipOutdated,
		SkipDeleted:  skipDeleted,
	}, utils.CtxErr{}
}
