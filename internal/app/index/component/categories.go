package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) categories() Node {
	return Div(
		htmx.Trigger("load"),
		htmx.Get(path.Categories(nil)),
		htmx.Swap("outerHTML"),
	)
}
