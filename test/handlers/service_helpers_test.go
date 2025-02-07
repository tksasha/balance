package handlers_test

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/db"
	"github.com/tksasha/balance/internal/handlers"
	providers "github.com/tksasha/balance/internal/providers/test"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/services"
)

func newCashService(ctx context.Context, t *testing.T) handlers.CashService {
	t.Helper()

	dbNameProvider := providers.NewDBNameProvider()

	db := db.Open(ctx, dbNameProvider)

	cashRepository := repositories.NewCashRepository(db)

	return services.NewCashService(cashRepository)
}
