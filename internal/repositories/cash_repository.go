package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
)

type CashRepository struct {
	db *sql.DB
}

func NewCashRepository(db *sql.DB) *CashRepository {
	return &CashRepository{
		db: db,
	}
}

func (r *CashRepository) Create(ctx context.Context, cash *models.Cash) error {
	currency := getCurrencyFromContext(ctx)

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

func (r *CashRepository) NameExists(ctx context.Context, name string, id int) (bool, error) {
	currency := getCurrencyFromContext(ctx)

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

func (r *CashRepository) FindByID(ctx context.Context, id int) (*models.Cash, error) {
	currency := getCurrencyFromContext(ctx)

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

	cash := &models.Cash{}

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

func (r *CashRepository) Update(ctx context.Context, cash *models.Cash) error {
	currency := getCurrencyFromContext(ctx)

	query := `
		UPDATE cashes
		SET
		    formula = ?,
		    sum = ?,
		    name = ?,
		    supercategory = ?,
		    favorite = ?
		WHERE
		    id = ?
		    AND currency = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		cash.Formula,
		cash.Sum,
		cash.Name,
		cash.Supercategory,
		cash.Favorite,
		cash.ID,
		currency,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrRecordNotFound
	}

	return nil
}

func (r *CashRepository) Delete(ctx context.Context, id int) error {
	currency := getCurrencyFromContext(ctx)

	query := `
		UPDATE cashes
		SET
			deleted_at = ?
		WHERE
			id = ?
			AND currency = ?
	`

	result, err := r.db.ExecContext(ctx, query, time.Now().UTC(), id, currency)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperrors.ErrRecordNotFound
	}

	return nil
}

func (r *CashRepository) List(ctx context.Context) (models.Cashes, error) {
	currency := getCurrencyFromContext(ctx)

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

	cashes := models.Cashes{}

	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		cash := &models.Cash{}

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
