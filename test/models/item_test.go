package model_test

import (
	"testing"
	"time"

	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSetDate(t *testing.T) {
	currencies := models.Currencies{}

	t.Run("when input value is blank it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetDate("")

		assert.Assert(t, !item.IsValid())
		assert.Assert(t, is.Contains(item.Errors.Get("date"), "is required"))
	})

	t.Run("when input value is invalid it should return an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetDate("2055")

		assert.Assert(t, !item.IsValid())
		assert.Assert(t, is.Contains(item.Errors.Get("date"), "is invalid"))
	})

	t.Run("when input value is valid it should set this value", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetDate("2024-09-30")

		date, err := time.Parse(time.DateOnly, "2024-09-30")
		if err != nil {
			panic(err)
		}

		assert.Equal(t, item.Date, date)
	})
}

func TestSetFormula(t *testing.T) {
	currencies := models.Currencies{}

	t.Run("when formula is blank it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetFormula("")

		assert.Assert(t, item.Errors.Has("formula"))
		assert.Assert(t, is.Contains(item.Errors.Get("formula"), "is required"))
	})

	t.Run("when formula is invalid it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetFormula("2/0")

		assert.Assert(t, item.Errors.Has("formula"))
		assert.Assert(t, is.Contains(item.Errors.Get("formula"), "is invalid"))
	})

	t.Run("when formula is valid it should set this value", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetFormula("10000+0.99")

		assert.Assert(t, !item.Errors.Has("formula"))
		assert.Equal(t, item.Formula, "10000+0.99")
		assert.Equal(t, item.Sum, 10_000.99)
	})
}

func TestSetCategoryID(t *testing.T) {
	currencies := models.Currencies{}

	t.Run("when input value is empty it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCategoryID("")

		assert.Assert(t, item.Errors.Has("category_id"))
		assert.Assert(t, is.Contains(item.Errors.Get("category_id"), "is required"))
	})

	t.Run("when input value is invalid it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCategoryID("abc")

		assert.Assert(t, item.Errors.Has("category_id"))
		assert.Assert(t, is.Contains(item.Errors.Get("category_id"), "is invalid"))
	})

	t.Run("when input value is valid it should set this value", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCategoryID("35")

		assert.Assert(t, !item.Errors.Has("category_id"))
		assert.Equal(t, item.CategoryID, 35)
	})
}

func TestSetDescription(t *testing.T) {
	currencies := models.Currencies{}

	item := models.NewItem(currencies)

	item.SetDescription("Rustic Wooden Keyboard")

	assert.Equal(t, item.Description, "Rustic Wooden Keyboard")
}

func TestSetCurrency(t *testing.T) {
	currencies := models.Currencies{
		"UAH": {ID: 0, Code: "UAH"},
		"USD": {ID: 1, Code: "USD"},
		"EUR": {ID: 3, Code: "EUR"},
	}

	t.Run("when currency is uah it should set uah", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCurrency("uah")

		assert.Equal(t, item.Currency.Code, "UAH")
	})

	t.Run("when currency is usd it should set usd", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCurrency("usd")

		assert.Equal(t, item.Currency.Code, "USD")
	})

	t.Run("when currency is eur it should set eur", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCurrency("eur")

		assert.Equal(t, item.Currency.Code, "EUR")
	})

	t.Run("when currency is empty it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCurrency("")

		assert.Assert(t, item.Errors.Has("currency"))
		assert.Assert(t, is.Contains(item.Errors.Get("currency"), "is required"))
	})

	t.Run("when currency is xxx it should set an error", func(t *testing.T) {
		item := models.NewItem(currencies)

		item.SetCurrency("xxx")

		assert.Assert(t, item.Errors.Has("currency"))
		assert.Assert(t, is.Contains(item.Errors.Get("currency"), "is invalid"))
	})
}
