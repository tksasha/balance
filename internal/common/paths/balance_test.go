package paths_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestBalance(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/balance?currency=uah", params: params.New()},
		{path: "/balance?currency=uah", params: params.New().WithCurrency(currency.UAH)},
		{path: "/balance?currency=usd", params: params.New().WithCurrency(currency.USD)},
		{path: "/balance?currency=eur", params: params.New().WithCurrency(currency.EUR)},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.Balance(d.params))
	}
}
