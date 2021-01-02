package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"
)

const (
	initUsersTable = `create table users(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
	username VARCHAR NOT NULL UNIQUE,
	password VARCHAR,
	full_name VARCHAR,
	PRIMARY KEY(id));`

	initQuestionsTable = `create table questions(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
	title VARCHAR NOT NULL,
	description TEXT NOT NULL,
	created_by BIGINT NOT NULL,
	created_at TIMESTAMP,
	status VARCHAR,
	rating INT,
	PRIMARY KEY(id),
	CONSTRAINT fk_question_user FOREIGN KEY (created_by) REFERENCES users(id));`

	initAnswersTable = `create table answers(
    id BIGINT GENERATED ALWAYS AS IDENTITY,
	question_id BIGINT NOT NULL,
	created_by BIGINT NOT NULL,
	created_at TIMESTAMP,
	answered BOOLEAN,
	PRIMARY KEY(id),
	CONSTRAINT fk_answer_user FOREIGN KEY (created_by) REFERENCES users(id),
	CONSTRAINT fk_answer_question FOREIGN KEY (question_id) REFERENCES questions(id));`
)

func init() {
	goose.AddMigration(Up20210101104725, Down20210101104725)
}

func Up20210101104725(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	ctx := context.Background()
	if _, err := tx.ExecContext(ctx, initUsersTable); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, initQuestionsTable); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, initAnswersTable); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func Down20210101104725(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
