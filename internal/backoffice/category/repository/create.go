package repository

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
)

func (r *Repository) Create(ctx context.Context, category *category.Category) error {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		INSERT INTO
		    categories (name, slug, income, visible, currency, supercategory, number)
		VALUES
		    (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Slug,
		category.Income,
		category.Visible,
		currency,
		category.Supercategory,
		category.Number,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = int(id)

	return nil
}
