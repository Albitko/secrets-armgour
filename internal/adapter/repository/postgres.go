package repository

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const schema = `
 	CREATE TABLE IF NOT EXISTS credentials_data (
 	    id serial primary key,
 		user_id text,
 		service text,
 		service_login text,
 		service_password text,
 		meta text,
 	    created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS text_data (
 	    id serial primary key,
 		user_id text,
 		title text,
 		note text,
 		meta text,
 	    created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS binary_data (
 	    id serial primary key,
 		user_id text,
 		title text,
 		data_content bytea,
 		meta text,
 	    created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS cards_data (
 	    id serial primary key,
 	    user_id text,
 		card_holder text,
 		card_number text,
 		card_validity_period text,
 		cvc_code text,
		meta text,
 	    created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS users_data (
 	    id serial primary key,
		user_login text not null unique,
 		password_hash text,
 	    created_at timestamp
 	);
`
const (
	uniqueViolationErr = "23505"
)

type logger interface {
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
}

type postgres struct {
	db *sql.DB
	l  logger
}

func (d *postgres) closeRows(row *sql.Rows) {
	err := row.Close()
	if err != nil {
		d.l.Errorf(
			"error closing rows %s", err.Error(),
		)
	}
}

func (d *postgres) queryRows(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var rows *sql.Rows

	statement, err := d.db.PrepareContext(
		ctx,
		query,
	)
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			d.l.Errorf(
				"error while closing statement %s", err.Error(),
			)
		}
	}(statement)

	if err != nil {
		d.l.Errorf(
			"error creating statement %s", err.Error(),
		)
		return rows, err
	}
	rows, err = statement.QueryContext(ctx, args...)

	if err != nil {
		d.l.Errorf(
			"error query execution %s", err.Error(),
		)
		return rows, err
	}

	if err = rows.Err(); err != nil {
		d.l.Errorf(
			"error row %s", err.Error(),
		)
		return rows, err
	}
	return rows, nil
}

func (d *postgres) closeStatement(statement *sql.Stmt) {
	if statement == nil {
		d.l.Errorf("error: nil statement")
		return
	}
	err := statement.Close()
	if err != nil {
		d.l.Errorf("error: %s closing statement", err.Error())
	}
}

func (d *postgres) Close() {
	err := d.db.Close()
	if err != nil {
		return
	}
}
func New(dsn string, l logger) (*postgres, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		l.Errorf("error: %s open postgres", err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		l.Errorf("error: %s ping postgres", err.Error())
		return nil, err
	}
	if _, err = db.Exec(schema); err != nil {
		l.Errorf("error: %s update schema", err.Error())
		return nil, err
	}
	return &postgres{
		l:  l,
		db: db,
	}, nil
}
