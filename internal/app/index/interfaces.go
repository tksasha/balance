package index

import "context"

type Repository interface {
	Income(ctx context.Context) (float64, error)
	Expense(ctx context.Context) (float64, error)
	Cashes(ctx context.Context) (float64, error)
}

type Service interface {
	Residual(ctx context.Context) (float64, error)
	Balance(ctx context.Context) (float64, error)
}
