package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/tksasha/balance/internal/apperrors"
)

type CashRepository struct {
	db *sql.DB
}

func NewCashRepository(db *sql.DB) *CashRepository {
	return &CashRepository{
		db: db,
	}
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
