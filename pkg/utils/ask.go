package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

// Confirm with a user
func Confirm(q string) (bool, CtxErr) {
	var confirmation bool
	prompt := &survey.Confirm{Message: q}

	err := survey.AskOne(prompt, &confirmation)
	return confirmation, CtxErr{
		Error:   err,
		Context: "Failed to get confirm",
	}
}
