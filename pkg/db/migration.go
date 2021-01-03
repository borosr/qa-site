package db

import (
	_ "github.com/borosr/qa-site/pkg/db/migartion"
	"github.com/pressly/goose"
)

func Migrate() error {
	return goose.Up(Get(), "pkg/db/migartion")
}
