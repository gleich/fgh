package configure

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

// Outline for the answers
type AnswersOutline struct {
	PAT      string
	Username string
}

// Ask questions to fill in configuration files
func AskQuestions() AnswersOutline {
	questions := []*survey.Question{
		{
			Name:   "PAT",
			Prompt: &survey.Password{Message: "What is your GitHub PAT?"},
		},
		{
			Name:   "Username",
			Prompt: &survey.Input{Message: "What is your GitHub username?"},
		},
	}
	var answers AnswersOutline
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask questions about config", err, 1)
	}
	return answers
}
