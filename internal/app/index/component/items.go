package component

import (
	"net/url"

	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) items() Node {
	return Div(Class("container"),
		Div(Class("row mt-4 mb-5"),
			Div(Class("col"),
				Div(Class("card items"),
					Div(Class("card-body"), ID("items"),
						htmx.Get(c.ListItems(0, 0, url.Values{})),
						htmx.Trigger("load"),
						Div(Class("spinner-border htmx-indicator"), ID("htmx-indicator")),
					),
				),
			),
		),
	)
}
