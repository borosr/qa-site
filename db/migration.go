package db

import "github.com/pressly/goose"

func Migrate() error {
	return goose.Up(Get(), "")
}
