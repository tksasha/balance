package paths_test

import (
	"testing"

	"github.com/tksasha/balance/internal/common/currency"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	"gotest.tools/v3/assert"
)

func TestCashes(t *testing.T) {
	ds := []struct {
		path   string
		params params.Params
	}{
		{path: "/cashes?currency=uah", params: params.New()},
		{path: "/cashes?currency=uah", params: params.New().WithCurrency(currency.UAH)},
		{path: "/cashes?currency=usd", params: params.New().WithCurrency(currency.USD)},
		{path: "/cashes?currency=eur", params: params.New().WithCurrency(currency.EUR)},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.Cashes(d.params))
	}
}

func TestEditCash(t *testing.T) {
	ds := []struct {
		id     int
		path   string
		params params.Params
	}{
		{id: 9, path: "/cashes/9/edit?currency=uah", params: params.New()},
		{id: 9, path: "/cashes/9/edit?currency=uah", params: params.New().WithCurrency(currency.UAH)},
		{id: 9, path: "/cashes/9/edit?currency=usd", params: params.New().WithCurrency(currency.USD)},
		{id: 9, path: "/cashes/9/edit?currency=eur", params: params.New().WithCurrency(currency.EUR)},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.EditCash(d.params, d.id))
	}
}

func TestUpdateCash(t *testing.T) {
	ds := []struct {
		id     int
		path   string
		params params.Params
	}{
		{id: 9, path: "/cashes/9?currency=uah", params: params.New()},
		{id: 9, path: "/cashes/9?currency=uah", params: params.New().WithCurrency(currency.UAH)},
		{id: 9, path: "/cashes/9?currency=usd", params: params.New().WithCurrency(currency.USD)},
		{id: 9, path: "/cashes/9?currency=eur", params: params.New().WithCurrency(currency.EUR)},
	}

	for _, d := range ds {
		assert.Equal(t, d.path, paths.UpdateCash(d.params, d.id))
	}
}
