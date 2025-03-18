package repository

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) Update(ctx context.Context, category *category.Category) error {
	query := `
		UPDATE
			categories
		SET
		    name = ?,
			slug = ?,
		    income = ?,
		    visible = ?,
		    supercategory = ?,
		    number = ?,
			currency = ?
		WHERE
		    id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		category.Name,
		category.Slug,
		category.Income,
		category.Visible,
		category.Supercategory,
		category.Number,
		category.Currency,
		category.ID,
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
