package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// FIXME remove this
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "postgres"
	dbname   = "qm-site"
)

var conn *sql.DB

func Get() *sql.DB {
	if conn == nil {
		var err error
		conn, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
		if err != nil {
			panic(err)
		}
	}

	return conn
}
