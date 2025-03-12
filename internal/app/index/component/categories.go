package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) categories(values url.Values) Node {
	return Div(
		htmx.Trigger("load"),
		htmx.Get(path.Categories(values, nil)),
		htmx.Swap("outerHTML"),
	)
}
