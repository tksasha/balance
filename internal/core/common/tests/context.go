package tests

import (
	"context"
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
)

func currencyContext(ctx context.Context, t *testing.T, curr currency.Currency) context.Context {
	t.Helper()

	switch curr {
	case currency.UAH:
		return context.WithValue(ctx, currency.ContextValue{}, currency.UAH)
	case currency.USD:
		return context.WithValue(ctx, currency.ContextValue{}, currency.USD)
	case currency.EUR:
		return context.WithValue(ctx, currency.ContextValue{}, currency.EUR)
	default:
		t.Fatalf("invalid currency: %v", curr)
	}

	return nil
}
