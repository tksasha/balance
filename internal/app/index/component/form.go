package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) form() Node {
	return Div(Class("container mt-4"),
		Div(Class("row"),
			Div(Class("col"),
				Div(Class("card"), ID("item-inline-form"),
					Div(Class("card-body"),
						htmx.Get(path.NewItem()),
						htmx.Trigger("load"),
					),
				),
			),
		),
	)
}
