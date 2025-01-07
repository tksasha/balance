package currencies_test

import (
	"testing"

	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestGetCurrencyIDByCode(t *testing.T) {
	data := map[string]int{
		"uah": 1,
		"UAH": 1,
		"usd": 2,
		"USD": 2,
		"eur": 3,
		"EUR": 3,
		"xxx": 0,
	}

	for code, id := range data {
		assert.Equal(t, currencies.GetCurrencyByCode(code), currencies.Currency(id))
	}
}

func TestGetDefaultCurrency(t *testing.T) {
	id, code := currencies.GetDefaultCurrency()

	assert.Equal(t, id, currencies.Currency(1))
	assert.Equal(t, code, "uah")
}
