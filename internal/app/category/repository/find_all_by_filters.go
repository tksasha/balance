package repository

import (
	"context"
	"log/slog"

	"github.com/tksasha/balance/internal/app/category"
)

func (r *Repository) FindAllByFilters(ctx context.Context, filters category.Filters) (category.Categories, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			items.category_id,
			items.category_name,
			SUM(items.sum),
			IIF(categories.income, 0, categories.supercategory) AS supercategory
		FROM
			items
		INNER JOIN
			categories
			ON categories.id = items.category_id
		WHERE
			items.deleted_at IS NULL
			AND items.currency = ?
			AND items.date BETWEEN ? AND ?
		GROUP BY
			items.category_id
	`

	rows, err := r.db.QueryContext(ctx, query, currency, filters.From, filters.To)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error("failed to close rows", "error", err)
		}
	}()

	categories := category.Categories{}

	for rows.Next() {
		category := &category.Category{}

		if err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Sum,
			&category.Supercategory,
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
