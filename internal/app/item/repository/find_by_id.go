package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) FindByID(ctx context.Context, id int) (*item.Item, error) {
	currency := r.GetCurrencyFromContext(ctx)

	query := `
		SELECT
		    id,
		    date,
			formula,
		    sum,
		    category_id,
		    category_name,
		    description
		FROM
		    items
		WHERE
		    items.id = ?
		    AND items.deleted_at IS NULL
		    AND currency = ?
	`

	item := &item.Item{}

	if err := r.db.
		QueryRowContext(ctx, query, id, currency).
		Scan(
			&item.ID,
			&item.Date,
			&item.Formula,
			&item.Sum,
			&item.CategoryID,
			&item.CategoryName,
			&item.Description,
		); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return item, nil
}
