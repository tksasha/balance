package repository

import (
	"context"

	"github.com/tksasha/balance/internal/common/repositories"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common"
)

func (r *Repository) Update(ctx context.Context, category *category.Category) error {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		UPDATE categories
		SET
		    name = ?,
		    income = ?,
		    visible = ?,
		    supercategory = ?
		WHERE
		    id = ?
		    AND currency = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Income,
		category.Visible,
		category.Supercategory,
		category.ID,
		currency,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.ErrRecordNotFound
	}

	return nil
}
