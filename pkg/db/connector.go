package db

import (
	"database/sql"

	"github.com/borosr/qa-site/pkg/healthcheck"
	"github.com/borosr/qa-site/pkg/settings"
	"github.com/dgraph-io/badger/v2"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var conn *sql.DB
var bdb *badger.DB

func Get() *sql.DB {
	if conn == nil {
		var err error
		conn, err = sql.Open("postgres", settings.Get().DBConnectionString)
		if err != nil {
			log.Error(err)
			healthcheck.Instance().Failed()
			return nil
		}
		healthcheck.Instance().Ok()
	}

	return conn
}

func GetBDB() *badger.DB {
	if bdb == nil {
		var err error
		bdb, err = badger.Open(badger.DefaultOptions(settings.Get().BadgerPath))
		if err != nil {
			log.Error(err)
			healthcheck.Instance().Failed()
			return nil
		}
		healthcheck.Instance().Ok()
	}

	return bdb
}
