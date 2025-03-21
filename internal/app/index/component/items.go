package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) items(params params.Params) Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4 mb-5"),
			Div(Class("col"),
				Div(Class("card items"),
					Div(Class("card-body"), ID("items"),
						htmx.Get(paths.Items(params)),
						htmx.Trigger("load"),
						Div(Class("spinner-border htmx-indicator"), ID("htmx-indicator")),
					),
				),
			),
		),
	)
}
