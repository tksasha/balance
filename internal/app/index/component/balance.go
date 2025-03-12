package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) balance() Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"),
			Div(
				htmx.Trigger("load"),
				htmx.Get(path.Balance(nil)),
				htmx.Swap("outerHTML"),
			),
			Div(
				htmx.Trigger("load"),
				htmx.Get(path.Cashes()),
				htmx.Swap("outerHTML"),
			),
		),
	)
}
