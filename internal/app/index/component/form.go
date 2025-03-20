package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form(params params.Params) Node {
	return Div(Class("container mt-4"),
		Div(Class("row"),
			Div(Class("col"),
				Div(Class("card"), ID("item-inline-form"),
					Div(Class("card-body"),
						htmx.Get(paths.NewItem(params)),
						htmx.Trigger("load"),
					),
				),
			),
		),
	)
}
