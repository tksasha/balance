package components

import (
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *ItemsComponent) Edit(item *item.Item, categories category.Categories) Node {
	return c.form(item, categories, nil)
}
