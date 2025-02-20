package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/pkg/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *ItemsComponent) Create(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	return c.form(item, categories, errors)
}
