package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/models"
)

type CashRepository struct {
	db *sql.DB
}

func NewCashRepository(db *db.DB) *CashRepository {
	return &CashRepository{
		db: db.Connection,
	}
}

func (r *CashRepository) Create(ctx context.Context, cash *models.Cash) error {
	currency := getCurrencyFromContext(ctx)

	query := `
		INSERT INTO cashes(currency, formula, sum, name, supercategory, favorite)
		VALUES(?, ?, ?, ?, ?, ?)
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
		SELECT 1
		FROM cashes
		WHERE
			currency = ? AND
			deleted_at IS NULL AND
			name = ? AND
			id != ?
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
		SELECT id, currency, name, formula, sum, supercategory, favorite
		FROM cashes
		WHERE currency = ? AND deleted_at IS NULL and id = ?
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
		SET formula = ?, sum = ?, name = ?, supercategory = ?, favorite = ?
		WHERE id = ? AND currency = ?
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
