package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) balance(params params.Params) Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"), ID("balance-and-cashes-row"),
			Div(
				htmx.Trigger("load"),
				htmx.Get(paths.Balance(params)),
				htmx.Swap("outerHTML"),
			),
			Div(
				htmx.Trigger("load"),
				htmx.Get(paths.Cashes(params)),
				htmx.Swap("outerHTML"),
			),
		),
	)
}
