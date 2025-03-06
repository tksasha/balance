package component

import (
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form() Node {
	return Div(Class("container mt-4"),
		Div(Class("row"),
			Div(Class("col"),
				Div(Class("card"),
					Div(Class("card-body"),
						htmx.Get("/items/new"),
						htmx.Trigger("load"),
						htmx.Swap("outerHTML"),
					),
				),
			),
		),
	)
}
