package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Update(params params.Params, cash *cash.Cash) Node {
	return c.ModalBody(
		c.Template(
			c.cash(params, cash, htmx.SwapOOB("true")),
		),
	)
}
