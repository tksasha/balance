package repository

import (
	"context"
	"log/slog"

	"github.com/tksasha/balance/internal/core/common/repositories"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/month"
)

func (r *Repository) List(ctx context.Context, month month.Month) (item.Items, error) {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			items.id,
			items.date,
			items.sum,
			COALESCE(items.category_name, ""),
			items.description
		FROM
			items
		WHERE
			items.currency = ?
		AND
			items.deleted_at IS NULL
		AND
			items.date BETWEEN ? AND ?
		ORDER BY
			items.date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, currency, month.Begin, month.End)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			slog.Error("failed to close rows", "error", err)
		}
	}()

	items := item.Items{}

	for rows.Next() {
		item := &item.Item{}

		if err := rows.Scan(&item.ID, &item.Date, &item.Sum, &item.CategoryName, &item.Description); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
