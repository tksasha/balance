package repository

import (
	"context"
)

func (r *Repository) Income(ctx context.Context) (float64, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    sum(items.sum)
		FROM
		    items
		    INNER JOIN categories ON categories.id = items.category_id
		WHERE
		    items.currency = ?
		    AND items.deleted_at IS NULL
		    AND categories.income = 1
	`

	row := r.db.QueryRowContext(ctx, query, currency)

	var income float64

	if err := row.Scan(&income); err != nil {
		return 0, err
	}

	return income, nil
}
