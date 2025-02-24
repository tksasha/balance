package repository

import (
	"context"
	"log"

	"github.com/tksasha/balance/internal/backoffice/category"
)

func (r *Repository) FindAll(ctx context.Context) (category.Categories, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			id,
			currency,
			name
		FROM
			categories
		WHERE
			currency = ?
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

	categories := category.Categories{}

	for rows.Next() {
		category := &category.Category{}

		if err := rows.Scan(&category.ID, &category.Currency, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
