package repository

import (
	"context"

	"github.com/tksasha/balance/internal/common/currency"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetCurrencyFromContext(ctx context.Context) currency.Currency {
	curr, ok := ctx.Value(currency.ContextValue{}).(currency.Currency)
	if !ok {
		curr = currency.Default
	}

	return curr
}
