package components

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validator"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *ItemsComponent) Create(item *item.Item, categories category.Categories, errors validator.Errors) Node {
	return c.form(item, categories, errors)
}
