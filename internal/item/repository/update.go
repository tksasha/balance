package repository

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/common/repositories"
	"github.com/tksasha/balance/internal/item"
)

func (r *Repository) Update(ctx context.Context, item *item.Item) error {
	currency := repositories.GetCurrencyFromContext(ctx)

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
		return apperrors.ErrRecordNotFound
	}

	return nil
}
