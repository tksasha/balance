package repository

import (
	"context"
	"log/slog"

	"github.com/tksasha/balance/internal/app/categoryreport"
)

func (r *Repository) Group(ctx context.Context, filters categoryreport.Filters) (categoryreport.Entities, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			items.category_name,
			items.category_slug,
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

	entities := categoryreport.Entities{}

	for rows.Next() {
		entity := &categoryreport.Entity{}

		if err := rows.Scan(
			&entity.CategoryName,
			&entity.CategorySlug,
			&entity.Sum,
			&entity.Supercategory,
		); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
