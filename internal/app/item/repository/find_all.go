package repository

import (
	"context"
	"log/slog"

	"github.com/tksasha/balance/internal/app/item"
)

func (r *Repository) FindAll(ctx context.Context, filters item.Filters) (item.Items, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			items.id,
			items.date,
			items.sum,
			items.category_name,
			items.description
		FROM
			items
		WHERE
			items.deleted_at IS NULL
			AND items.currency = ?
			AND items.date between ? AND ?
	`

	args := []any{currency, filters.From, filters.To}

	if filters.Category != 0 {
		query += ` AND items.category_id = ?`

		args = append(args, filters.Category)
	}

	query += ` ORDER BY date DESC`

	rows, err := r.db.QueryContext(ctx, query, args...)
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
