package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"
)

const (
	initRevokeTokensTable = `create table if not exists revoke_tokens(
    id VARCHAR,
	owner_id VARCHAR NOT NULL,
	token VARCHAR NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_revoke_token_user FOREIGN KEY (owner_id) REFERENCES users(id));`
)

func init() {
	goose.AddMigration(Up20210103015836, Down20210103015836)
}

func Up20210103015836(tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(ctx, initRevokeTokensTable); err != nil {
		return err
	}
	return nil
}

func Down20210103015836(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
