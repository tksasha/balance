package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/repositories"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List(ctx context.Context) (cash.Cashes, error) {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			id,
			currency,
			formula,
			sum,
			name,
			supercategory,
			favorite
		FROM
			cashes
		WHERE
			currency = ?
			AND deleted_at IS NULL
	`

	rows, err := r.db.QueryContext(ctx, query, currency)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	cashes := cash.Cashes{}

	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		cash := &cash.Cash{}

		if err := rows.Scan(
			&cash.ID,
			&cash.Currency,
			&cash.Formula,
			&cash.Sum,
			&cash.Name,
			&cash.Supercategory,
			&cash.Favorite,
		); err != nil {
			return nil, err
		}

		cashes = append(cashes, cash)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cashes, nil
}
