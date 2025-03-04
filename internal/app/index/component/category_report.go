package component

import (
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) categoryReport() Node {
	return Div(
		htmx.Trigger("load"),
		htmx.Get("/categoryreport"),
		htmx.Swap("outerHTML"),
	)
}
