package utils

import (
	"goforge/internal/models"

	"github.com/AlecAivazis/survey/v2"
)

func AskQuestions(questions []*models.Question) (models.Answers, error) {
	answers := make(models.Answers)

	for _, q := range questions {
		if err := ask(q, answers); err != nil {
			return nil, err
		}
	}

	return answers, nil
}

func ask(q *models.Question, answers models.Answers) error {

	
	if len(q.Options) == 0 { //Then response is a confirm (yes/no)

		var value bool

		err := survey.AskOne(&survey.Confirm{
				Message: q.Text,
			},
			&value,
		)
		if err != nil {
			return err
		}

		answers[q.Key] = value

		if value {
			if q.Next != nil {
				if err := ask(q.Next, answers); err != nil {
					return err
				}
			}
		}

		return nil
	}

	// MultiSelect
	var values []string

	err := survey.AskOne(&survey.MultiSelect{
			Message: q.Text,
			Options: q.Options,
		},
		&values,
	)
	if err != nil {
		return err
	}

	answers[q.Key] = values

	return nil
}