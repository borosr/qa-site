package questions

import (
	"github.com/borosr/qa-site/pkg/questions/repository"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	// TODO add tags
}

type PageableResponse struct {
	Data  []repository.QuestionWithRating `json:"data"`
	Count int64                           `json:"count"`
}
