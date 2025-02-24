package common

import (
	"context"

	"github.com/tksasha/balance/internal/common/currency"
)

type BaseRepository struct{}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{}
}

func (r *BaseRepository) GetCurrencyFromContext(ctx context.Context) currency.Currency {
	curr, ok := ctx.Value(currency.ContextValue{}).(currency.Currency)
	if !ok {
		curr = currency.Default
	}

	return curr
}
