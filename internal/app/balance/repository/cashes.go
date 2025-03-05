package repository

import (
	"context"
)

func (r *Repository) Cashes(ctx context.Context) (float64, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			COALESCE(SUM(sum), 0)
		FROM
			cashes
		WHERE
			currency = ?
			AND deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, currency)

	var cashes float64

	if err := row.Scan(&cashes); err != nil {
		return 0, err
	}

	return cashes, nil
}
