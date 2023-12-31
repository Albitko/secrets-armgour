package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Albitko/secrets-armgour/internal/entity"
)

// DeleteUserData - remove users secrets
func (d *postgres) DeleteUserData(ctx context.Context, data, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	var query string
	switch data {
	case entity.Credentials:
		query = `delete from credentials_data where id = $1`
	case entity.Binary:
		query = `delete from binary_data where id = $1`
	case entity.Text:
		query = `delete from text_data where id = $1`
	case entity.Card:
		query = `delete from cards_data where id = $1`
	default:
		d.l.Errorf(
			"error unsupported data type",
		)
		cancel()
		return fmt.Errorf("unsupported data type")
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
