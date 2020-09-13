package configure

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

type AnswersOutline struct {
	PAT string
}

// Ask questions to fill in configuration files
func AskQuestions() AnswersOutline {
	questions := []*survey.Question{
		{
			Name:   "PAT",
			Prompt: &survey.Password{Message: "What is your GitHub PAT?"},
		},
	}
	var answers AnswersOutline
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask questions about config", err, 1)
	}

	fmt.Println(answers)
	return answers
}
