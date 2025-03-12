package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) balance(values url.Values) Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"),
			Div(
				htmx.Trigger("load"),
				htmx.Get(path.Balance(values)),
				htmx.Swap("outerHTML"),
			),
			Div(
				htmx.Trigger("load"),
				htmx.Get(path.Cashes(values)),
				htmx.Swap("outerHTML"),
			),
		),
	)
}
