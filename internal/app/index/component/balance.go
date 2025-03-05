package component

import (
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) balance() Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"),
			Div(
				htmx.Trigger("load"),
				htmx.Get("/balance"),
				htmx.Swap("outerHTML"),
			),
			Div(
				htmx.Trigger("load"),
				htmx.Get("/cashes"),
				htmx.Swap("outerHTML"),
			),
		),
	)
}
