package models_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/pkg/currencies"
	"gotest.tools/v3/assert"
)

func TestItemBuilder_WithDate(t *testing.T) {
	t.Run("when input is blank it should return error", func(t *testing.T) {
		_, err := models.NewItemBuilder().WithDate("").Build()

		assert.Equal(t, err.String(), "date: is required, is invalid")
	})

	t.Run("when input is invalid it should return error", func(t *testing.T) {
		_, err := models.NewItemBuilder().WithDate("abc").Build()

		assert.Equal(t, err.String(), "date: is invalid")
	})

	t.Run("when input is valid it should return item with date", func(t *testing.T) {
		item, validationErrors := models.NewItemBuilder().WithDate("2024-10-15").Build()

		date, err := time.Parse(time.DateOnly, "2024-10-15")
		if err != nil {
			panic(err)
		}

		assert.Equal(t, item.Date, date)
		assert.Assert(t, validationErrors.IsEmpty())
	})
}

func TestItemBuilder_WithFormula(t *testing.T) {
	t.Run("when input is blank it should set an error", func(t *testing.T) {
		_, err := models.NewItemBuilder().WithFormula("").Build()

		assert.Equal(t, err.String(), "formula: is required, is invalid")
	})

	t.Run("when input is invalid it should set an error", func(t *testing.T) {
		_, err := models.NewItemBuilder().WithFormula("2/0").Build()

		assert.Equal(t, err.String(), "formula: is invalid")
	})

	t.Run("when input is valid it should calculate sum", func(t *testing.T) {
		item, validationErrors := models.NewItemBuilder().WithFormula("10000+0.99").Build()

		assert.Equal(t, item.Formula, "10000+0.99")
		assert.Equal(t, item.Sum, 10_000.99)
		assert.Assert(t, validationErrors.IsEmpty())
	})
}

func TestItemBuilder_WithCategoryID(t *testing.T) {
	t.Run("when input value is empty it should set an error", func(t *testing.T) {
		_, err := models.NewItemBuilder().WithCategoryID("").Build()

		assert.Equal(t, err.String(), "category_id: is required, is invalid")
	})

	t.Run("when input value is invalid it should set an error", func(t *testing.T) {
		_, err := models.NewItemBuilder().WithCategoryID("abc").Build()

		assert.Equal(t, err.String(), "category_id: is invalid")
	})

	t.Run("when input value is valid it should set this value", func(t *testing.T) {
		item, validationErrors := models.NewItemBuilder().WithCategoryID("35").Build()

		assert.Equal(t, item.CategoryID, 35)
		assert.Assert(t, validationErrors.IsEmpty())
	})
}

func TestItemBuilder_WithDescription(t *testing.T) {
	item, err := models.NewItemBuilder().WithDescription("Rustic Wooden Keyboard").Build()

	assert.Equal(t, item.Description, "Rustic Wooden Keyboard")
	assert.Assert(t, err.IsEmpty())

	t.Run("when input contains leading or trailing spaces it should trim them", func(t *testing.T) {
		item, err := models.NewItemBuilder().WithDescription(" Stellar Breeze Catalyst ").Build()

		assert.Equal(t, item.Description, "Stellar Breeze Catalyst")
		assert.Assert(t, err.IsEmpty())
	})
}

func TestItemBuilder_WithCurrencyCode(t *testing.T) {
	t.Run("when code is uah it should set id to 1", func(t *testing.T) {
		item, err := models.NewItemBuilder().WithCurrencyCode("uah").Build()

		assert.Equal(t, item.Currency, currencies.Currency(1))
		assert.Equal(t, item.CurrencyCode, "uah")
		assert.Assert(t, err.IsEmpty())
	})

	t.Run("when code is usd it should set id to 2", func(t *testing.T) {
		item, err := models.NewItemBuilder().WithCurrencyCode("usd").Build()

		assert.Equal(t, item.Currency, currencies.Currency(2))
		assert.Equal(t, item.CurrencyCode, "usd")
		assert.Assert(t, err.IsEmpty())
	})

	t.Run("when code is eur it should set id to 3", func(t *testing.T) {
		item, err := models.NewItemBuilder().WithCurrencyCode("eur").Build()

		assert.Equal(t, item.Currency, currencies.Currency(3))
		assert.Equal(t, item.CurrencyCode, "eur")
		assert.Assert(t, err.IsEmpty())
	})

	t.Run("when code is xxx it should set default currency", func(t *testing.T) {
		item, err := models.NewItemBuilder().WithCurrencyCode("xxx").Build()

		assert.Equal(t, item.Currency, currencies.UAH)
		assert.Assert(t, err.IsEmpty())
	})
}
