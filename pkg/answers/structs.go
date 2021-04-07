package answers

type Request struct {
	UpdateRequest
	QuestionID string `json:"question_id"`
}

type UpdateRequest struct {
	Answer string `json:"answer"`
}
