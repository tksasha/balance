package repository

import (
	"context"

	"github.com/tksasha/balance/internal/core/common/repositories"
)

func (r *Repository) Cashes(ctx context.Context) (float64, error) {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			sum(sum)
		FROM
			cashes
		WHERE
			currency = ?
			AND deleted_at IS NULL
	`

	row := r.db.QueryRowContext(ctx, query, currency)

	var cashes float64

	if err := row.Scan(&cashes); err != nil {
		return 0.0, err
	}

	return cashes, nil
}
