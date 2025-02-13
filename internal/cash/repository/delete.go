package repository

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/repositories"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	currency := repositories.GetCurrencyFromContext(ctx)

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
