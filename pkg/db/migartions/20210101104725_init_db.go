package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose"
)

const (
	initUsersTable = `create table if not exists users(
    id VARCHAR,
	username VARCHAR NOT NULL UNIQUE,
	password VARCHAR,
	full_name VARCHAR,
	PRIMARY KEY(id));`

	initQuestionsTable = `create table if not exists questions(
    id VARCHAR,
	title VARCHAR NOT NULL,
	description TEXT NOT NULL,
	created_by VARCHAR NOT NULL,
	created_at TIMESTAMP,
	status VARCHAR,
	rating INT,
	PRIMARY KEY(id),
	CONSTRAINT fk_question_user FOREIGN KEY (created_by) REFERENCES users(id));`

	initAnswersTable = `create table if not exists answers(
    id VARCHAR,
	question_id VARCHAR NOT NULL,
	created_by VARCHAR NOT NULL,
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
	ctx := context.Background()
	if _, err := tx.ExecContext(ctx, initUsersTable); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, initQuestionsTable); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, initAnswersTable); err != nil {
		return err
	}

	return nil
}

func Down20210101104725(tx *sql.Tx) error {
	return nil
}
