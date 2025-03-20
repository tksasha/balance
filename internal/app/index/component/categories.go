package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) categories(params params.Params) Node {
	return Div(
		htmx.Trigger("load"),
		htmx.Get(paths.Categories(params)),
		htmx.Swap("outerHTML"),
	)
}
