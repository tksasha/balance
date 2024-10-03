package helpers_test

import (
	"testing"

	"github.com/tksasha/balance/internal/helpers"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
)

func TestEditItemURL(t *testing.T) {
	item := models.Item{
		ID: 1409,
	}

	assert.Equal(t, helpers.EditItemURL(item), "/items/1409/edit")
}
