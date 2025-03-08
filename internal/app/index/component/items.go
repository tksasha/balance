package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
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
						htmx.Get(path.Items(nil, nil)),
						htmx.Trigger("load"),
						Div(Class("spinner-border htmx-indicator"), ID("htmx-indicator")),
					),
				),
			),
		),
	)
}
