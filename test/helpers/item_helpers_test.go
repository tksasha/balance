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
	decorator := decorators.NewItemDecorator(
		&models.Item{
			ID: 1409,
		},
	)

	res := helpers.EditItemURL(decorator)

	exp := templ.URL("/items/1409/edit")

	assert.Equal(t, res, exp)
}
