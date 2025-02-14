package repository

import (
	"context"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/common/repositories"
)

func (r *Repository) Create(ctx context.Context, category *category.Category) error {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		INSERT INTO
		    categories (name, income, visible, currency, supercategory)
		VALUES
		    (?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Income,
		category.Visible,
		currency,
		category.Supercategory,
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
