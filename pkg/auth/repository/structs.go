package repository

import "time"

type TokenCache struct {
	Token string
	Expr  time.Time
}
