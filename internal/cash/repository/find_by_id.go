package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/repositories"
)

func (r *Repository) FindByID(ctx context.Context, id int) (*cash.Cash, error) {
	currency := repositories.GetCurrencyFromContext(ctx)

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
		&cash.Favorite,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrRecordNotFound
		}

		return nil, err
	}

	return cash, nil
}
