package repository

import (
	"context"

	"github.com/tksasha/balance/internal/backoffice/cash"
)

func (r *Repository) Create(ctx context.Context, cash *cash.Cash) error {
	query := `
		INSERT INTO
		    cashes (
		        currency,
		        formula,
		        sum,
		        name,
		        supercategory
		    )
		VALUES
		    (?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		cash.Currency,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	cash.ID = int(id)

	return nil
}
