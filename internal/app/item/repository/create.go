package repository

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/app/item"
)

func (r *Repository) Create(ctx context.Context, item *item.Item) error {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		INSERT INTO items (
			date,
			formula,
			sum,
			category_id,
			category_name,
			category_slug,
			description,
			currency
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		item.Date.Format(time.DateOnly),
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.CategorySlug,
		item.Description,
		currency,
	); err != nil {
		return err
	}

	return nil
}
