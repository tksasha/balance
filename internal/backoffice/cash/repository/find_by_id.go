package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) FindByID(ctx context.Context, id int) (*cash.Cash, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    id,
		    currency,
		    name,
		    formula,
		    sum,
		    supercategory,
		    favorite
		FROM
		    cashes
		WHERE
		    deleted_at IS NULL
		    AND id = ?
			AND currency = ?
	`

	row := r.db.QueryRowContext(ctx, query, id, currency)

	cash := &cash.Cash{}

	if err := row.Scan(
		&cash.ID,
		&cash.Currency,
		&cash.Name,
		&cash.Formula,
		&cash.Sum,
		&cash.Supercategory,
		&cash.Favorite,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return cash, nil
}
