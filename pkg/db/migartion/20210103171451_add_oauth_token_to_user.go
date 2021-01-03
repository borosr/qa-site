package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"
)

const (
	alterUsersWithToken = `alter table users add access_token VARCHAR;`
)

func init() {
	goose.AddMigration(Up20210103171451, Down20210103171451)
}

func Up20210103171451(tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(ctx, alterUsersWithToken); err != nil {
		return err
	}

	return nil
}

func Down20210103171451(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
