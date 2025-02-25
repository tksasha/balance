package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/app/cash"
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
		    supercategory
		FROM
		    cashes
		WHERE
		    currency = ?
		    AND deleted_at IS NULL
		    AND id = ?
	`

	row := r.db.QueryRowContext(ctx, query, currency, id)

	cash := &cash.Cash{}

	if err := row.Scan(
		&cash.ID,
		&cash.Currency,
		&cash.Name,
		&cash.Formula,
		&cash.Sum,
		&cash.Supercategory,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return cash, nil
}
