package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) linkToBackoffice() Node {
	return Div(Class("link-to-backoffice"), Title("backoffice"),
		htmx.Get(paths.Backoffice()),
		htmx.Target("#modal-body"),
		htmx.SwapOOB("true"),
		Data("bs-toggle", "modal"),
		Data("bs-target", "#modal"),
		Text("B"),
	)
}
