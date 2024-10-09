package helpers_test

import (
	"testing"

	"github.com/a-h/templ"
	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/helpers"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestEditItemURL(t *testing.T) {
	currency := models.Currency{
		ID:   3,
		Name: "eur",
	}

	decorator := decorators.NewItemDecorator(
		&models.Item{
			ID: 1409,
		},
	)

	res := helpers.EditItemURL(currency, decorator)

	exp := templ.URL("/eur/items/1409/edit")

	assert.Equal(t, res, exp)
}
