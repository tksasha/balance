package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Update(cash *cash.Cash) Node {
	return c.ModalBody(
		c.Template(
			c.cash(cash, htmx.SwapOOB("true")),
		),
	)
}
