package decorators_test

import (
	"testing"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNewItemsDecorator(t *testing.T) {
	item := decorators.NewItemDecorator(&models.Item{})

	decorator := decorators.NewItemsDecorator(
		[]*models.Item{
			item.Item,
		},
	)

	assert.Assert(t, is.Contains(decorator.Items, item))
}
