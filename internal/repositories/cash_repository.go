package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/db"
	internalerrors "github.com/tksasha/balance/internal/errors"
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

func (r *CashRepository) FindByName(ctx context.Context, name string) (*models.Cash, error) {
	currency := getCurrencyFromContext(ctx)

	query := `
		SELECT id, currency, name, formula, sum, supercategory, favorite
		FROM cashes
		WHERE currency = ? AND deleted_at IS NULL and name = ?
	`

	row := r.db.QueryRowContext(ctx, query, currency, name)

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
			return nil, internalerrors.ErrRecordNotFound
		}

		return nil, err
	}

	return cash, nil
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
