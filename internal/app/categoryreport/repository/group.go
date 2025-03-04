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
			category_name,
			category_slug,
			SUM(items.sum)
		FROM
			items
		WHERE
			deleted_at IS NULL
			AND currency = ?
			AND date BETWEEN ? AND ?
		GROUP BY
			category_name
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

		if err := rows.Scan(&entity.CategoryName, &entity.CategorySlug, &entity.Sum); err != nil {
			return nil, err
		}

		entities = append(entities, entity)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entities, nil
}
