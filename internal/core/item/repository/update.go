package repository

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/core/common"
	"github.com/tksasha/balance/internal/core/item"
)

func (r *Repository) Update(ctx context.Context, item *item.Item) error {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		UPDATE items
		SET
		    date = ?,
		    formula = ?,
		    sum = ?,
		    category_id = ?,
		    category_name = ?,
		    description = ?
		WHERE
		    id = ?
		    AND deleted_at IS NULL
		    AND currency = ?
	`

	result, err := r.db.ExecContext(ctx, query,
		item.Date.Format(time.DateOnly),
		item.Formula,
		item.Sum,
		item.CategoryID,
		item.CategoryName,
		item.Description,
		item.ID,
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
