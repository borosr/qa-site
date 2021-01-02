package db

import (
	_ "github.com/borosr/qa-site/pkg/db/migartions"
	"github.com/pressly/goose"
)

func Migrate() error {
	return goose.Up(Get(), "pkg/db/migartions")
}
