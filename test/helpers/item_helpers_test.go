package helpers_test

import (
	"testing"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/helpers"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestItemsURL(t *testing.T) {
	t.Run("when currency is provided it should return url with currency", func(t *testing.T) {
		currency := "eur"

		res := helpers.ItemsURL(&currency)

		exp := templ.URL("/items?currency=eur")

		assert.Equal(t, res, exp)
	})

	t.Run("when currency is not provided it should return url without currency", func(t *testing.T) {
		res := helpers.ItemsURL(nil)

		exp := templ.URL("/items")

		assert.Equal(t, res, exp)
	})
}

func TestItemURL(t *testing.T) {
	item := &models.Item{
		ID: 1331,
	}

	res := helpers.ItemURL(item)

	exp := templ.URL("/items/1331")

	assert.Equal(t, res, exp)
}

func TestEditItemURL(t *testing.T) {
	item := &models.Item{
		ID: 1409,
	}

	res := helpers.EditItemURL(item)

	exp := templ.URL("/items/1409/edit")

	assert.Equal(t, res, exp)
}
