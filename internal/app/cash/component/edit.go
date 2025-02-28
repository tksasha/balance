package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/validation"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
)

func (c *Component) Edit(cash *cash.Cash, errors validation.Errors) Node {
	nodes := []Node{
		htmx.Patch(c.cashUpdatePath(cash.ID)),
		htmx.Target("#modal-body"),
		htmx.Swap("outerHTML"),
	}

	return c.form(
		Map(nodes, func(node Node) Node { return node }),
		cash,
		errors,
	)
}
