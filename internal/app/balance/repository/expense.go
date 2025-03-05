package repository

import (
	"context"

	"github.com/shopspring/decimal"
)

func (r *Repository) Expense(ctx context.Context) (decimal.Decimal, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    COALESCE(SUM(items.sum), 0)
		FROM
		    items
		    INNER JOIN categories ON categories.id = items.category_id
		WHERE
		    items.currency = ?
		    AND items.deleted_at IS NULL
		    AND categories.income = 0
	`

	row := r.db.QueryRowContext(ctx, query, currency)

	var income decimal.Decimal

	if err := row.Scan(&income); err != nil {
		return decimal.NewFromInt(0), err
	}

	return income, nil
}
