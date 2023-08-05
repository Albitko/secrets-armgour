package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

func (d *postgres) UpdateCard(ctx context.Context, index int, card entity.UserCard) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateCard, err := d.db.PrepareContext(
		ctx, "UPDATE cards_data SET card_holder=$1, card_number=$2, card_validity_period=$3, cvc_code=$4, meta=$5 WHERE id=$6;",
	)
	if err != nil {
		d.l.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			d.l.Errorf(
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
		d.l.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) UpdateCredentials(ctx context.Context, index int, credentials entity.UserCredentials) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateCredentials, err := d.db.PrepareContext(
		ctx, "UPDATE credentials_data SET service=$1, service_login=$2, service_password=$3, meta=$4 WHERE id=$5;",
	)
	if err != nil {
		d.l.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			d.l.Errorf(
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
		d.l.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) UpdateBinary(ctx context.Context, index int, bin entity.UserBinary, data []byte) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateBinary, err := d.db.PrepareContext(
		ctx, "UPDATE binary_data SET title=$1, data_content=$2, meta=$3 WHERE id=$4;",
	)
	if err != nil {
		d.l.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			d.l.Errorf(
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
		d.l.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}

func (d *postgres) UpdateText(ctx context.Context, index int, text entity.UserText) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updateText, err := d.db.PrepareContext(
		ctx, "UPDATE text_data SET title=$1, note=$2, meta=$3 WHERE id=$4;",
	)
	if err != nil {
		d.l.Errorf(
			"error creating statement %s", err.Error(),
		)
		return err
	}
	defer func(updateUser *sql.Stmt) {
		err := updateUser.Close()
		if err != nil {
			d.l.Errorf(
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
		d.l.Errorf(
			"error executing query %s", err.Error(),
		)
		return err
	}
	return nil
}
