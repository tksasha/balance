package repository

import (
	"context"

	"github.com/shopspring/decimal"
)

func (r *Repository) Cashes(ctx context.Context) (decimal.Decimal, error) {
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

	var cashes decimal.Decimal

	if err := row.Scan(&cashes); err != nil {
		return decimal.NewFromInt(0), err
	}

	return cashes, nil
}
