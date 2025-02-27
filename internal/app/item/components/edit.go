package components

import (
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *ItemsComponent) Edit(item *item.Item, categories category.Categories, errors validation.Errors) Node {
	return c.form(item, categories, errors)
}
