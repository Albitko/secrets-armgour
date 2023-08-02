package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Albitko/secrets-armgour/internal/utils/logger"
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
 		card_holder text,
 		card_number text,
 		card_validity_period text,
 		cvc_code smallint,
		meta text,
 	    created_at timestamp
 	);
	CREATE TABLE IF NOT EXISTS users (
 	    id serial primary key,
 		user_id text,
 		password_hash text,
 	    created_at timestamp
 	);
`

type postgres struct {
	db *sql.DB
}

func (d *postgres) Close() {
	err := d.db.Close()
	if err != nil {
		return
	}
}
func New(psqlConn string) (*postgres, error) {
	db, err := sql.Open("pgx", psqlConn)
	if err != nil {
		logger.Zap.Errorf("error: %s open postgres", err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logger.Zap.Errorf("error: %s ping postgres", err.Error())
		return nil, err
	}
	if _, err = db.Exec(schema); err != nil {
		logger.Zap.Errorf("error: %s update schema", err.Error())
		return nil, err
	}
	return &postgres{
		db: db,
	}, nil
}
