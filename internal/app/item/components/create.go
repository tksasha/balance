package components

import (
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *ItemsComponent) Create(item *item.Item) Node {
	return c.item(item)
}
