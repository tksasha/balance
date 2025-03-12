package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Update(values url.Values, item *item.Item) Node {
	return c.Template(c.item(values, item, htmx.SwapOOB("true")))
}
