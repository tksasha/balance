package component

import (
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
)

func (c *Component) Update(item *item.Item) Node {
	return c.item(item)
}
