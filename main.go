package main

import (
	"time"

	"github.com/borosr/qa-site/internal/api"
	"github.com/borosr/qa-site/pkg/db"
	"github.com/borosr/qa-site/pkg/healthcheck"
	log "github.com/sirupsen/logrus"
)

func main() {
	initLog()

	if healthcheck.Instance().Healthy() {
		if err := db.Migrate(); err != nil {
			healthcheck.Instance().Failed()
			log.Error(err)
		}
	}
	if err := api.Init(); err != nil {
		log.Error(err)
	}
}

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  time.RFC3339,
		QuoteEmptyFields: true,
	})
}
