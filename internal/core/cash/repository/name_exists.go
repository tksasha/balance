package repository

import (
	"context"
	"database/sql"
	"errors"
)

func (r *Repository) NameExists(ctx context.Context, name string, id int) (bool, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			1
		FROM
			cashes
		WHERE
			currency = ?
			AND deleted_at IS NULL
			AND name = ?
			AND id != ?
	`

	row := r.db.QueryRowContext(ctx, query, currency, name, id)

	var exists int

	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return exists == 1, nil
}
