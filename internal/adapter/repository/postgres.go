package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Albitko/secrets-armgour/internal/entity"
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
 	    user_id text,
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

func closeStatement(statement *sql.Stmt) {
	if statement == nil {
		logger.Zap.Errorf("error: nil statement")
		return
	}
	err := statement.Close()
	if err != nil {
		logger.Zap.Errorf("error: %s closing statement", err.Error())
	}
}

func (d *postgres) InsertCard(card entity.UserCard) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertCard, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO cards_data (card_holder, card_number, card_validity_period, cvc_code, meta, created_at) VALUES ($1, $2, $3, $4, $5, $6);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertCard)

	_, err = insertCard.ExecContext(
		ctx,
		card.CardHolder,
		card.CardNumber,
		card.CardValidityPeriod,
		card.CvcCode,
		card.Meta,
		createdAt,
	)
	if err != nil {
		logger.Zap.Errorf("error: %s write card data", err.Error())
		return err
	}
	return nil
}

func (d *postgres) InsertCredentials(credentials entity.UserCredentials) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertCredentials, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO credentials_data (service, service_login, service_password,  meta, created_at) VALUES ($1, $2, $3, $4, $5);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertCredentials)

	_, err = insertCredentials.ExecContext(
		ctx,
		credentials.ServiceName,
		credentials.ServiceLogin,
		credentials.ServicePassword,
		credentials.Meta,
		createdAt,
	)
	if err != nil {
		logger.Zap.Errorf("error: %s write credentials data", err.Error())
		return err
	}
	return nil
}

func (d *postgres) InsertBinary(bin entity.UserBinary, data []byte) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertBinary, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO binary_data (title, data_content, meta, created_at) VALUES ($1, $2, $3, $4);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertBinary)

	_, err = insertBinary.ExecContext(
		ctx,
		bin.Title,
		data,
		bin.Meta,
		createdAt,
	)
	if err != nil {
		logger.Zap.Errorf("error: %s write binary data", err.Error())
		return err
	}
	return nil
}

func (d *postgres) InsertText(text entity.UserText) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertText, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO text_data (title, note, meta, created_at) VALUES ($1, $2, $3, $4);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertText)

	_, err = insertText.ExecContext(
		ctx,
		text.Title,
		text.Body,
		text.Meta,
		createdAt,
	)
	if err != nil {
		logger.Zap.Errorf("error: %s write text data", err.Error())
		return err
	}
	return nil
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
