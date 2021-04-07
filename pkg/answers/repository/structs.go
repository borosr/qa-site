package repository

import "github.com/volatiletech/null/v8"

type AnswerWithRating struct {
	ID         string    `boil:"id" json:"id"`
	QuestionID string    `boil:"question_id" json:"question_id"`
	CreatedBy  string    `boil:"created_by" json:"created_by"`
	Answer     string    `boil:"answer" json:"answer"`
	CreatedAt  null.Time `boil:"created_at" json:"created_at,omitempty"`
	Answered   null.Bool `boil:"answered" json:"answered,omitempty"`
	Rating     int       `boil:"rating" json:"rating"`
}
