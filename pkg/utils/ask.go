package utils

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

func Confirm(q string) bool {
	var confirmation bool
	prompt := &survey.Confirm{Message: q}

	err := survey.AskOne(prompt, &confirmation)
	if err != nil {
		statuser.Error("Failed to confirm", err, 1)
	}
	return confirmation
}
