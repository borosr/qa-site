package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"
)

const (
	addRating = `create table if not exists ratings(
    id VARCHAR,
    kind VARCHAR NOT NULL,
    record_id VARCHAR NOT NULL,
    rated_by VARCHAR NOT NULL,
    rated_at TIMESTAMP NOT NULL,
    value SMALLINT NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_rating_user FOREIGN KEY (rated_by) REFERENCES users(id)
	);`
)

func init() {
	goose.AddMigration(Up20210105094435, Down20210105094435)
}

func Up20210105094435(tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(ctx, addRating); err != nil {
		return err
	}

	return nil
}

func Down20210105094435(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
