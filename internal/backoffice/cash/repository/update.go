package repository

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) Update(ctx context.Context, cash *cash.Cash) error {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		UPDATE cashes
		SET
		    formula = ?,
		    sum = ?,
		    name = ?,
		    supercategory = ?,
		    favorite = ?
		WHERE
		    id = ?
		    AND currency = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
		cash.Favorite,
		cash.ID,
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
