package utils

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/handlers"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
)

func NewCashService(ctx context.Context, t *testing.T) (handlers.CashService, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	cashRepository := repositories.NewCashRepository(db)

	return services.NewCashService(cashRepository), db
}
