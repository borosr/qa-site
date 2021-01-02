package main

import (
	"log"

	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/src/api"
)

func main() {
	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	if err := api.Init(); err != nil {
		log.Fatal(err)
	}
}
