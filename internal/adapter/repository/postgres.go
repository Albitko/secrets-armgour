package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
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

type postgres struct {
	db *sql.DB
}

func (d *postgres) GetUserPasswordHash(login string) (string, error) {
	var passHash string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	getSecret, err := d.db.PrepareContext(
		ctx, `select password_hash from users_data where user_login = $1;`,
	)
	if err != nil {
		logger.Zap.Errorf(
			"error prepare context %s", err.Error(),
		)
		return "", err
	}
	defer getSecret.Close()

	err = getSecret.QueryRowContext(ctx, login).Scan(&passHash)
	if err != nil {
		logger.Zap.Errorf(
			"error query execution %s", err.Error(),
		)
		return "", err
	}
	return passHash, nil
}

func (d *postgres) RegisterUser(login, pass string) error {
	var pgErr *pgconn.PgError
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertCard, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO users_data (user_login, password_hash, created_at) VALUES ($1, $2, $3);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertCard)

	_, err = insertCard.ExecContext(
		ctx,
		login,
		pass,
		createdAt,
	)
	if err != nil && errors.As(err, &pgErr) {
		if pgErr.Code == uniqueViolationErr {
			logger.Zap.Info("login in use")
			return entity.ErrLoginAlreadyInUse
		} else {
			logger.Zap.Errorf("error: %s write user auth data", err.Error())
			return err
		}
	}
	return nil
}

func (d *postgres) UpdateCard(index int, card entity.UserCard) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateCard, err := d.db.PrepareContext(
		ctx, "UPDATE cards_data SET card_holder=$1, card_number=$2, card_validity_period=$3, cvc_code=$4, meta=$5 WHERE id=$6;",
	)
	if err != nil {
		logger.Zap.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			logger.Zap.Errorf(
				"error closing statement %s", err.Error(),
			)
		}
	}(updateCard)

	_, err = updateCard.ExecContext(
		ctx,
		card.CardHolder,
		card.CardNumber,
		card.CardValidityPeriod,
		card.CvcCode,
		card.Meta,
		index,
	)
	if err != nil {
		logger.Zap.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) UpdateCredentials(index int, credentials entity.UserCredentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateCredentials, err := d.db.PrepareContext(
		ctx, "UPDATE credentials_data SET service=$1, service_login=$2, service_password=$3, meta=$4 WHERE id=$5;",
	)
	if err != nil {
		logger.Zap.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			logger.Zap.Errorf(
				"error closing statement %s", err.Error(),
			)
		}
	}(updateCredentials)

	_, err = updateCredentials.ExecContext(
		ctx,
		credentials.ServiceName,
		credentials.ServiceLogin,
		credentials.ServicePassword,
		credentials.Meta,
		index,
	)
	if err != nil {
		logger.Zap.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) UpdateBinary(index int, bin entity.UserBinary, data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateBinary, err := d.db.PrepareContext(
		ctx, "UPDATE binary_data SET title=$1, data_content=$2, meta=$3 WHERE id=$4;",
	)
	if err != nil {
		logger.Zap.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			logger.Zap.Errorf(
				"error closing statement %s", err.Error(),
			)
		}
	}(updateBinary)

	_, err = updateBinary.ExecContext(
		ctx,
		bin.Title,
		data,
		bin.Meta,
		index,
	)
	if err != nil {
		logger.Zap.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) UpdateText(index int, text entity.UserText) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updateText, err := d.db.PrepareContext(
		ctx, "UPDATE text_data SET title=$1, note=$2, meta=$3 WHERE id=$4;",
	)
	if err != nil {
		logger.Zap.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			logger.Zap.Errorf(
				"error closing statement %s", err.Error(),
			)
		}
	}(updateText)

	_, err = updateText.ExecContext(
		ctx,
		text.Title,
		text.Body,
		text.Meta,
		index,
	)
	if err != nil {
		logger.Zap.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) DeleteUserData(data, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var query string
	switch data {
	case "credentials":
		query = `delete from credentials_data where id = $1`
	case "binary":
		query = `delete from binary_data where id = $1`
	case "text":
		query = `delete from text_data where id = $1`
	case "card":
		query = `delete from cards_data where id = $1`
	}
	defer cancel()
	delStmnt, err := d.db.PrepareContext(
		ctx,
		query,
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(delStmnt)
	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Zap.Errorf("error: %s converting string to int", err.Error())
		return err
	}
	_, err = delStmnt.ExecContext(
		ctx,
		intId,
	)
	if err != nil {
		logger.Zap.Errorf("error: %s write text data", err.Error())
		return err
	}
	return nil
}

func (d *postgres) GetUserData(data, id, user string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var res interface{}

	switch data {
	case "credentials":
		var credential entity.UserCredentials
		query := `
		select service,service_login,service_password,meta  from credentials_data where id = $1 and user_id = $2;
		`
		getSecret, err := d.db.PrepareContext(
			ctx, query,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error prepare context %s", err.Error(),
			)
			return "", err
		}
		defer getSecret.Close()

		err = getSecret.QueryRowContext(ctx, id, user).Scan(
			&credential.ServiceName,
			&credential.ServiceLogin,
			&credential.ServicePassword,
			&credential.Meta,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return "", err
		}
		res = credential
	case "binary":
		var bin entity.UserBinary
		query := `
		select title,data_content,meta  from binary_data where id = $1 and user_id = $2;
		`
		getSecret, err := d.db.PrepareContext(
			ctx, query,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error prepare context %s", err.Error(),
			)
			return "", err
		}
		defer getSecret.Close()

		err = getSecret.QueryRowContext(ctx, id, user).Scan(
			&bin.Title,
			&bin.B64Content,
			&bin.Meta,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return "", err
		}
		res = bin

	case "text":
		var text entity.UserText
		query := `
		select title,note,meta  from text_data where id = $1 and user_id = $2;
		`
		getSecret, err := d.db.PrepareContext(
			ctx, query,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error prepare context %s", err.Error(),
			)
			return "", err
		}
		defer getSecret.Close()

		err = getSecret.QueryRowContext(ctx, id, user).Scan(
			&text.Title,
			&text.Body,
			&text.Meta,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return "", err
		}
		res = text
	case "card":
		var card entity.UserCard
		query := `
		select card_holder,card_number,card_validity_period,cvc_code,meta  from cards_data where id = $1 and user_id = $2;
		`
		getSecret, err := d.db.PrepareContext(
			ctx, query,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error prepare context %s", err.Error(),
			)
			return "", err
		}
		defer getSecret.Close()

		err = getSecret.QueryRowContext(ctx, id, user).Scan(
			&card.CardHolder,
			&card.CardNumber,
			&card.CardValidityPeriod,
			&card.CvcCode,
			&card.Meta,
		)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return "", err
		}
		res = card
	}

	return res, nil
}

func closeRows(row *sql.Rows) {
	err := row.Close()
	if err != nil {
		logger.Zap.Errorf(
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
			logger.Zap.Errorf(
				"error while closing statement %s", err.Error(),
			)
		}
	}(statement)

	if err != nil {
		logger.Zap.Errorf(
			"error creating statement %s", err.Error(),
		)
		return rows, err
	}
	rows, err = statement.QueryContext(ctx, args...)

	if err != nil {
		logger.Zap.Errorf(
			"error query execution %s", err.Error(),
		)
		return rows, err
	}

	if err = rows.Err(); err != nil {
		logger.Zap.Errorf(
			"error row %s", err.Error(),
		)
		return rows, err
	}
	return rows, nil
}

func (d *postgres) SelectUserData(data, user string) (interface{}, error) {
	var rows *sql.Rows
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var res interface{}

	switch data {
	case "credentials":
		var credentials []entity.CutCredentials
		var credential entity.CutCredentials
		query := `
		select id,service,meta  from credentials_data where user_id = $1;
		`
		rows, err = d.queryRows(ctx, query, user)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return res, err
		}
		for rows.Next() {
			err = rows.Scan(&credential.Id, &credential.ServiceName, &credential.Meta)
			if err != nil {
				return res, err
			}
			credentials = append(credentials, credential)
		}
		res = credentials
	case "binary":
		var bins []entity.CutBinary
		var bin entity.CutBinary
		query := `
		select id,title,meta  from binary_data where user_id = $1;
		`
		rows, err = d.queryRows(ctx, query, user)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return res, err
		}
		for rows.Next() {
			err = rows.Scan(&bin.Id, &bin.Title, &bin.Meta)
			if err != nil {
				return res, err
			}
			bins = append(bins, bin)
		}
		res = bins
	case "text":
		var texts []entity.CutText
		var text entity.CutText
		query := `
		select id,title,meta  from text_data where user_id = $1;
		`
		rows, err = d.queryRows(ctx, query, user)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return res, err
		}
		for rows.Next() {
			err = rows.Scan(&text.Id, &text.Title, &text.Meta)
			if err != nil {
				return res, err
			}
			texts = append(texts, text)
		}
		fmt.Println(texts)
		res = texts
	case "card":
		var cards []entity.CutCard
		var card entity.CutCard
		query := `
		select id,card_number,meta  from cards_data where user_id = $1;
		`
		rows, err = d.queryRows(ctx, query, user)
		if err != nil {
			logger.Zap.Errorf(
				"error query execution %s", err.Error(),
			)
			return res, err
		}
		for rows.Next() {
			err = rows.Scan(&card.Id, &card.CardNumber, &card.Meta)
			if err != nil {
				return res, err
			}
			cards = append(cards, card)
		}
		res = cards
	}
	defer closeRows(rows)

	return res, nil
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

func (d *postgres) InsertCard(card entity.UserCard, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertCard, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO cards_data (user_id,card_holder, card_number, card_validity_period, cvc_code, meta, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertCard)

	_, err = insertCard.ExecContext(
		ctx,
		user,
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

func (d *postgres) InsertCredentials(credentials entity.UserCredentials, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertCredentials, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO credentials_data (user_id,service, service_login, service_password,  meta, created_at) VALUES ($1, $2, $3, $4, $5, $6);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertCredentials)

	_, err = insertCredentials.ExecContext(
		ctx,
		user,
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

func (d *postgres) InsertBinary(bin entity.UserBinary, data []byte, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertBinary, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO binary_data (user_id,title, data_content, meta, created_at) VALUES ($1, $2, $3, $4, $5);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertBinary)

	_, err = insertBinary.ExecContext(
		ctx,
		user,
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

func (d *postgres) InsertText(text entity.UserText, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	insertText, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO text_data (user_id,title, note, meta, created_at) VALUES ($1, $2, $3, $4, $5);",
	)
	if err != nil {
		logger.Zap.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer closeStatement(insertText)

	_, err = insertText.ExecContext(
		ctx,
		user,
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
