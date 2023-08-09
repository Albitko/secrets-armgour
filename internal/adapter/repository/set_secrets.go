package repository

import (
	"context"
	"time"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

// InsertCard - save card data for user
func (d *postgres) InsertCard(ctx context.Context, card entity.UserCard, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	insertCard, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO cards_data (user_id,card_holder, card_number, card_validity_period, cvc_code, meta, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);",
	)
	if err != nil {
		d.l.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer d.closeStatement(insertCard)

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
		d.l.Errorf("error: %s write card data", err.Error())
		return err
	}
	return nil
}

// InsertCredentials - save credentials data for user
func (d *postgres) InsertCredentials(ctx context.Context, credentials entity.UserCredentials, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	insertCredentials, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO credentials_data (user_id,service, service_login, service_password,  meta, created_at) VALUES ($1, $2, $3, $4, $5, $6);",
	)
	if err != nil {
		d.l.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer d.closeStatement(insertCredentials)

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
		d.l.Errorf("error: %s write credentials data", err.Error())
		return err
	}
	return nil
}

// InsertBinary - save credentials data for user
func (d *postgres) InsertBinary(ctx context.Context, bin entity.UserBinary, data []byte, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	insertBinary, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO binary_data (user_id,title, data_content, meta, created_at) VALUES ($1, $2, $3, $4, $5);",
	)
	if err != nil {
		d.l.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer d.closeStatement(insertBinary)

	_, err = insertBinary.ExecContext(
		ctx,
		user,
		bin.Title,
		data,
		bin.Meta,
		createdAt,
	)
	if err != nil {
		d.l.Errorf("error: %s write binary data", err.Error())
		return err
	}
	return nil
}

// InsertText - save text data for user
func (d *postgres) InsertText(ctx context.Context, text entity.UserText, user string) error {
	now := time.Now()
	createdAt := now.Format("2006-01-02T15:04")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	insertText, err := d.db.PrepareContext(
		ctx,
		"INSERT INTO text_data (user_id,title, note, meta, created_at) VALUES ($1, $2, $3, $4, $5);",
	)
	if err != nil {
		d.l.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer d.closeStatement(insertText)

	_, err = insertText.ExecContext(
		ctx,
		user,
		text.Title,
		text.Body,
		text.Meta,
		createdAt,
	)
	if err != nil {
		d.l.Errorf("error: %s write text data", err.Error())
		return err
	}
	return nil
}
