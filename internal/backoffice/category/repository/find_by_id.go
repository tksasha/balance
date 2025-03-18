package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) FindByID(ctx context.Context, id int) (*category.Category, error) {
	query := `
		SELECT
		    id,
		    name,
		    income,
		    visible,
		    currency,
		    supercategory,
			number
		FROM
		    categories
		WHERE
		    id = ?
	`

	row := r.db.QueryRowContext(ctx, query, id)

	category := &category.Category{}

	if err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Income,
		&category.Visible,
		&category.Currency,
		&category.Supercategory,
		&category.Number,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.ErrRecordNotFound
		}

		return nil, err
	}

	return category, nil
}
