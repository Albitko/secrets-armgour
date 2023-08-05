package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

func (d *postgres) GetUserData(ctx context.Context, data, id, user string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
				"error query execution %s", err.Error(),
			)
			return "", err
		}
		res = card
	}

	return res, nil
}

func (d *postgres) SelectUserData(ctx context.Context, data, user string) (interface{}, error) {
	var rows *sql.Rows
	var err error
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
			d.l.Errorf(
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
	defer d.closeRows(rows)

	return res, nil
}
