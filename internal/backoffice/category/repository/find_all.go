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
			name,
			supercategory,
			visible,
			income,
			number
		FROM
			categories
		WHERE
			currency = ?
		ORDER BY
			number
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

		if err := rows.Scan(
			&category.ID,
			&category.Currency,
			&category.Name,
			&category.Supercategory,
			&category.Visible,
			&category.Income,
			&category.Number,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
