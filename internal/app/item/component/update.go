package component

import (
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Update(item *item.Item) Node {
	return c.Template(c.item(item, htmx.SwapOOB("true")))
}
