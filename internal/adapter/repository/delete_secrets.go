package repository

import (
	"context"
	"strconv"
	"time"
)

func (d *postgres) DeleteUserData(ctx context.Context, data, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
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
		d.l.Errorf("error: %s preparing statement", err.Error())
		return err
	}
	defer d.closeStatement(delStmnt)
	intId, err := strconv.Atoi(id)
	if err != nil {
		d.l.Errorf("error: %s converting string to int", err.Error())
		return err
	}
	_, err = delStmnt.ExecContext(
		ctx,
		intId,
	)
	if err != nil {
		d.l.Errorf("error: %s write text data", err.Error())
		return err
	}
	return nil
}
