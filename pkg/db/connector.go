package db

import (
	"database/sql"

	"github.com/borosr/qa-site/pkg/settings"
	_ "github.com/lib/pq"
)

var conn *sql.DB

func Get() *sql.DB {
	if conn == nil {
		var err error
		conn, err = sql.Open("postgres", settings.Get().DBConnectionString)
		if err != nil {
			panic(err)
		}
	}

	return conn
}
