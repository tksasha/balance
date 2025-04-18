package balance

import (
	"context"
)

type Repository interface {
	Income(ctx context.Context) (float64, error)
	Expense(ctx context.Context) (float64, error)
	Cashes(ctx context.Context) (float64, error)
}

type Service interface {
	Balance(ctx context.Context) (*Balance, error)
}
