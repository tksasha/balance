package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) linkToBackoffice() Node {
	return Div(Class("link-to-backoffice"), Title("backoffice"),
		htmx.Get(path.Backoffice()),
		htmx.Target("#backoffice-modal-body"),
		htmx.SwapOOB("true"),
		Data("bs-toggle", "modal"),
		Data("bs-target", "#backoffice-modal"),
		Text("B"),
	)
}
