package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) FindByName(ctx context.Context, name string) (*category.Category, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    id,
		    name,
		    currency
		FROM
		    categories
		WHERE
		    name = ?
		    AND currency = ?
		LIMIT
		    1
	`

	category := &category.Category{}

	row := r.db.QueryRowContext(ctx, query, name, currency)

	if err := row.Scan(&category.ID, &category.Name, &category.Currency); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}
