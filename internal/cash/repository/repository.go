package repository

import (
	"context"
	"database/sql"
	"errors"
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

func (r *Repository) NameExists(ctx context.Context, name string, id int) (bool, error) {
	currency := repositories.GetCurrencyFromContext(ctx)

	query := `
		SELECT
			1
		FROM
			cashes
		WHERE
			currency = ?
			AND deleted_at IS NULL
			AND name = ?
			AND id != ?
	`

	row := r.db.QueryRowContext(ctx, query, currency, name, id)

	var exists int

	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return exists == 1, nil
}

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
