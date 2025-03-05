package balance

import (
	"context"

	"github.com/shopspring/decimal"
)

type Repository interface {
	Income(ctx context.Context) (decimal.Decimal, error)
	Expense(ctx context.Context) (decimal.Decimal, error)
	Cashes(ctx context.Context) (decimal.Decimal, error)
}

type Service interface {
	Balance(ctx context.Context) (*Balance, error)
}
