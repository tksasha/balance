package model_test

import (
	"testing"

	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSetDate(t *testing.T) {
	t.Run("when input value is blank it should set an error", func(t *testing.T) {
		item := models.NewItem()

		item.SetDate("")

		assert.Assert(t, !item.IsValid())
		assert.Assert(t, is.Contains(item.Errors.Get("date"), "is required"))
	})

	t.Run("when input value is invalid it should return an error", func(t *testing.T) {
		item := models.NewItem()

		item.SetDate("2055")

		assert.Assert(t, !item.IsValid())
		assert.Assert(t, is.Contains(item.Errors.Get("date"), "is invalid"))
	})

	t.Run("when input value is valid it should set this value", func(t *testing.T) {
		item := models.NewItem()

		item.SetDate("2024-09-30")

		assert.Equal(t, item.GetDateAsString(), "2024-09-30")
	})
}

func TestSetFormula(t *testing.T) {
	t.Run("when formula is blank it should set an error", func(t *testing.T) {
		item := models.NewItem()

		item.SetFormula("")

		assert.Assert(t, item.Errors.Has("formula"))
		assert.Assert(t, is.Contains(item.Errors.Get("formula"), "is required"))
	})

	t.Run("when formula is invalid it should set an error", func(t *testing.T) {
		item := models.NewItem()

		item.SetFormula("2/0")

		assert.Assert(t, item.Errors.Has("formula"))
		assert.Assert(t, is.Contains(item.Errors.Get("formula"), "is invalid"))
	})

	t.Run("when formula is valid it should set this value", func(t *testing.T) {
		item := models.NewItem()

		item.SetFormula("10000+0.99")

		assert.Assert(t, !item.Errors.Has("formula"))
		assert.Equal(t, item.Formula, "10000+0.99")
		assert.Equal(t, item.Sum, 10_000.99)
		assert.Equal(t, item.GetSumAsString(), "10\u00A0000,99")
	})
}

func TestSetCategoryID(t *testing.T) {
	t.Run("when input value is empty it should set an error", func(t *testing.T) {
		item := models.NewItem()

		item.SetCategoryID("")

		assert.Assert(t, item.Errors.Has("category_id"))
		assert.Assert(t, is.Contains(item.Errors.Get("category_id"), "is required"))
	})

	t.Run("when input value is invalid it should set an error", func(t *testing.T) {
		item := models.NewItem()

		item.SetCategoryID("abc")

		assert.Assert(t, item.Errors.Has("category_id"))
		assert.Assert(t, is.Contains(item.Errors.Get("category_id"), "is invalid"))
	})

	t.Run("when input value is valid it should set this value", func(t *testing.T) {
		item := models.NewItem()

		item.SetCategoryID("35")

		assert.Assert(t, !item.Errors.Has("category_id"))
		assert.Equal(t, item.CategoryID, 35)
		assert.Equal(t, item.GetCategoryIDAsString(), "35")
	})
}

func TestSetDescription(t *testing.T) {
	item := models.NewItem()

	item.SetDescription("Rustic Wooden Keyboard")

	assert.Equal(t, item.GetDescription(), "Rustic Wooden Keyboard")
}
