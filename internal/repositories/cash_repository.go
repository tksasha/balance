package repositories

import (
	"context"
	"database/sql"
	"errors"
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

func (r *CashRepository) NameExists(ctx context.Context, name string, id int) (bool, error) { // TODO: delme
	return false, nil
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
