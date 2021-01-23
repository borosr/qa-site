package questions

import (
	"github.com/volatiletech/null/v8"
)

type Request struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	// TODO add tags
}

type PageableResponse struct {
	Data  []Response `json:"data"`
	Count int64      `json:"count"`
}

type Response struct {
	ID          string      `boil:"id" json:"id"`
	Title       string      `boil:"title" json:"title"`
	Description string      `boil:"description" json:"description"`
	CreatedBy   string      `boil:"created_by" json:"created_by"`
	CreatedAt   null.Time   `boil:"created_at" json:"created_at,omitempty"`
	Status      null.String `boil:"status" json:"status,omitempty"`
	Rating      int         `json:"rating" boil:"rating"`
}
