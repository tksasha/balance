package index

import (
	"context"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/category"
)

type Repository interface {
	Income(ctx context.Context) (float64, error)
	Expense(ctx context.Context) (float64, error)
	Cashes(ctx context.Context) (float64, error)
}

type Service interface {
	Residual(ctx context.Context) (float64, error)
	Balance(ctx context.Context) (float64, error)
}

type CategoryService interface {
	List(ctx context.Context) (category.Categories, error)
}

type CashService interface {
	List(ctx context.Context) (cash.Cashes, error)
}
