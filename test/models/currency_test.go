package models_test

import (
	"testing"

	"github.com/tksasha/balance/internal/models"
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
		assert.Equal(t, models.GetCurrencyByCode(code), models.Currency(id))
	}
}

func TestGetDefaultCurrency(t *testing.T) {
	id, code := models.GetDefaultCurrency()

	assert.Equal(t, id, models.Currency(1))
	assert.Equal(t, code, "uah")
}
