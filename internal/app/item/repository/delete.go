package repository

import (
	"context"
	"time"

	"github.com/tksasha/balance/internal/common"
)

func (r *Repository) Delete(ctx context.Context, id int) error {
	currencies := r.GetCurrencyFromContext(ctx)

	query := `
		UPDATE items
		SET
		    deleted_at = ?
		WHERE
		    id = ?
		    AND deleted_at IS NULL
		    AND currency = ?
	`

	result, err := r.db.ExecContext(ctx, query, time.Now().UTC(), id, currencies)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.ErrRecordNotFound
	}

	return nil
}
