package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/cash"
)

func (r *Repository) NameExists(ctx context.Context, cash *cash.Cash) (bool, error) {
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

	row := r.db.QueryRowContext(ctx, query, cash.Currency, cash.Name, cash.ID)

	var exists int

	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return exists == 1, nil
}
