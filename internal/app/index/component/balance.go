package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) balance(cashes cash.Cashes) Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"),
			Div(
				htmx.Trigger("load"),
				htmx.Get("/balance"),
				htmx.Swap("outerHTML"),
			),
			c.cashComponent.List(cashes),
		),
	)
}
