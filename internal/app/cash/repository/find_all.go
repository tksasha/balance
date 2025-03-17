package repository

import (
	"context"
	"log/slog"

	"github.com/tksasha/balance/internal/app/cash"
)

func (r *Repository) FindAll(ctx context.Context) (cash.Cashes, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			id,
			name,
			sum,
			supercategory
		FROM
			cashes
		WHERE
			deleted_at IS NULL
			AND currency = ?
	`

	rows, err := r.db.QueryContext(ctx, query, currency)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error("failed to close rows", "error", err)
		}
	}()

	cashes := cash.Cashes{}

	for rows.Next() {
		cash := &cash.Cash{}

		if err := rows.Scan(&cash.ID, &cash.Name, &cash.Sum, &cash.Supercategory); err != nil {
			return nil, err
		}

		cashes = append(cashes, cash)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cashes, nil
}
