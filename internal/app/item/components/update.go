package components

import (
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *ItemsComponent) Update(item *item.Item) Node {
	return c.item(item)
}
