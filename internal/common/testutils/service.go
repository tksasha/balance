package testutils

import (
	"context"
	"database/sql"
	"testing"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/cash/repository"
	"github.com/tksasha/balance/internal/cash/service"
)

func NewCashService(ctx context.Context, t *testing.T) (cash.Service, *sql.DB) {
	t.Helper()

	db := newDB(ctx, t)

	cashRepository := repository.New(db)

	return service.New(cashRepository), db
}
