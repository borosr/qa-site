package main

import (
	"log"

	"github.com/borosr/qa-site/src/api"
)

func main() {
	if err := api.Init(); err != nil {
		log.Fatal(err)
	}
}
