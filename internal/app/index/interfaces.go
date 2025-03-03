package index

import (
	"context"

	"github.com/tksasha/balance/internal/app/balance"
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
	Balance(ctx context.Context) (*balance.Balance, error)
}
