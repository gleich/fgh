package configure

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

type RegularOutline struct {
	CloneClipboard bool `yaml:"clone_clipboard"`
}

type SecretsOutline struct {
	PAT string `yaml:"pat"`
}

// Ask questions to fill in reglar config
func AskQuestions() RegularOutline {
	questions := []*survey.Question{
		{
			Name:   "CloneClipboard",
			Prompt: &survey.Confirm{Message: "Do you want to copy the path of a cloned repo after clone to your clipboard?"},
		},
	}
	var answers RegularOutline
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask questions about config", err, 1)
	}
	return answers
}

// Ask questions to fill in the secret config
func AskSecretQuestions() SecretsOutline {
	questions := []*survey.Question{
		{
			Name: "PAT",
			Prompt: &survey.Input{
				Message: "What is your GitHub PAT?",
				Help:    "Get a token from https://github.com/settings/tokens/new with the repo box checked",
			},
		},
	}
	var answers SecretsOutline
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask questions about config", err, 1)
	}
	return answers
}
