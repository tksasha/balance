package repository

import (
	"context"
	"log"

	"github.com/tksasha/balance/internal/backoffice/cash"
)

func (r *Repository) FindAll(ctx context.Context) (cash.Cashes, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			id,
			currency,
			formula,
			sum,
			name,
			supercategory
		FROM
			cashes
		WHERE
			currency = ?
			AND deleted_at IS NULL
		ORDER BY
			name ASC
	`

	rows, err := r.db.QueryContext(ctx, query, currency)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	cashes := cash.Cashes{}

	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		cash := &cash.Cash{}

		if err := rows.Scan(
			&cash.ID,
			&cash.Currency,
			&cash.Formula,
			&cash.Sum,
			&cash.Name,
			&cash.Supercategory,
		); err != nil {
			return nil, err
		}

		cashes = append(cashes, cash)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cashes, nil
}
