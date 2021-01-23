package ratings

const (
	QuestionKind kind = "question"
	AnswerKind   kind = "answer"
)

type kind string

type Response struct {
	Value int64 `json:"value"`
}
