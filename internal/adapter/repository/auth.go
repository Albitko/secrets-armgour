package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

// GetUserPasswordHash - return password hash for login
func (d *postgres) GetUserPasswordHash(ctx context.Context, login string) (string, error) {
	var passHash string
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	getSecret, err := d.db.PrepareContext(
		ctx, `select password_hash from users_data where user_login = $1;`,
	)
	if err != nil {
		d.l.Errorf(
			"error prepare context %s", err.Error(),
		)
		return "", err
	}
	defer getSecret.Close()

	err = getSecret.QueryRowContext(ctx, login).Scan(&passHash)
	if err != nil {
		d.l.Errorf(
			"error query execution %s", err.Error(),
		)
		return "", err
	}
	return passHash, nil
}

// RegisterUser - save login and password pair for service
func (d *postgres) RegisterUser(ctx context.Context, login, pass string) error {
	var pgErr *pgconn.PgError
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	insertCard, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO users_data (user_login, password_hash, created_at) VALUES ($1, $2, $3);",
	)
	if err != nil {
		d.l.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer d.closeStatement(insertCard)

	_, err = insertCard.ExecContext(
		ctx,
		login,
		pass,
		createdAt,
	)
	if err != nil && errors.As(err, &pgErr) {
		if pgErr.Code == uniqueViolationErr {
			d.l.Infof("login in use")
			return entity.ErrLoginAlreadyInUse
		} else {
			d.l.Errorf("error: %s write user auth data", err.Error())
			return err
		}
	}
	return nil
}
