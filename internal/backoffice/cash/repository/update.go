package repository

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) Update(ctx context.Context, cash *cash.Cash) error {
	query := `
		UPDATE cashes
		SET
		    formula = ?,
		    sum = ?,
		    name = ?,
		    supercategory = ?,
			currency = ?
		WHERE
		    id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
		cash.Currency,
		cash.ID,
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
