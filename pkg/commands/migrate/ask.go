package migrate

import (
	"fmt"
	"os"

	"github.com/Matt-Gleich/fgh/pkg/utils"
)

// Confirm with the user that they want to move the repos
func ConfirmMove(newPaths map[string]string) utils.CtxErr {
	move, err := utils.Confirm(
		fmt.Sprintf("Are you sure you want to move %v repos?", len(newPaths)),
	)
	if err.Error != nil {
		return err
	}

	if !move {
		os.Exit(0)
	}
	return utils.CtxErr{}
}
