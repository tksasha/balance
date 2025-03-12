package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Update(values url.Values, cash *cash.Cash) Node {
	return c.ModalBody(
		c.Template(
			c.cash(values, cash, htmx.SwapOOB("true")),
		),
	)
}
