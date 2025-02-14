package repository

import (
	"context"

	"github.com/tksasha/balance/internal/common/repositories"
	"github.com/tksasha/balance/internal/core/cash"
)

func (r *Repository) Create(ctx context.Context, cash *cash.Cash) error {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		INSERT INTO
		    cashes (
		        currency,
		        formula,
		        sum,
		        name,
		        supercategory,
		        favorite
		    )
		VALUES
		    (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		currency,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
		cash.Favorite,
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
