package main

import (
	"time"

	"github.com/borosr/qa-site/internall/api"
	"github.com/borosr/qa-site/pkg/db"
	log "github.com/sirupsen/logrus"
)

func main() {
	initLog()

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}
	if err := api.Init(); err != nil {
		log.Fatal(err)
	}
}

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:             true,
		TimestampFormat:           time.RFC3339,
		QuoteEmptyFields:          true,
	})
}
