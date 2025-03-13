package path_test

import (
	"maps"
	"testing"

	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/balance/internal/common/currency"
	"gotest.tools/v3/assert"
)

func TestNewCurrency(t *testing.T) {
	ds := []struct {
		currency currency.Currency
		params   path.Params
	}{
		{
			currency: currency.UAH,
			params:   path.Params{"currency": "uah"},
		},
		{
			currency: currency.USD,
			params:   path.Params{"currency": "usd"},
		},
		{
			currency: currency.EUR,
			params:   path.Params{"currency": "eur"},
		},
	}

	for _, d := range ds {
		expected := d.params

		actual := path.NewCurrency(d.currency)

		assert.Assert(t, maps.Equal(expected, actual))
	}
}
