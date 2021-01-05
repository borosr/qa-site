package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"
)

const (
	addAnswerColumnToAnswerTable = `alter table answers add answer TEXT NOT NULL;`
	removeRatingColumnFromQuestionsTable = `alter table questions drop column rating;`
)

func init() {
	goose.AddMigration(Up20210103215215, Down20210103215215)
}

func Up20210103215215(tx *sql.Tx) error {
	ctx := context.Background()
	if _, err := tx.ExecContext(ctx, addAnswerColumnToAnswerTable); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, removeRatingColumnFromQuestionsTable); err != nil {
		return err
	}

	return nil
}

func Down20210103215215(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
