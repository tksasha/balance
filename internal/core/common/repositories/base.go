package repositories

import (
	"context"

	"github.com/tksasha/balance/pkg/currencies"
)

type BaseRepository struct{}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{}
}

func (r *BaseRepository) GetCurrencyFromContext(ctx context.Context) currencies.Currency {
	currency, ok := ctx.Value(currencies.CurrencyContextValue{}).(currencies.Currency)
	if !ok {
		currency = currencies.DefaultCurrency
	}

	return currency
}
