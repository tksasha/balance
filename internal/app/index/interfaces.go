package index

import (
	"context"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/category"
)

type CategoryService interface {
	List(ctx context.Context) (category.Categories, error)
}

type CashService interface {
	List(ctx context.Context) (cash.Cashes, error)
}

type BalanceService interface {
	Balance(ctx context.Context) (float64, float64, error)
}
