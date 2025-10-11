package component

import (
	"github.com/tksasha/balance/internal/common/paths"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Breadcrumbs(children ...Node) Node {
	return Nav(Class("mb-4"),
		Ol(Class("breadcrumb"),
			Li(Class("breadcrumb-item"),
				Span(Class("link"), Text("Backoffice"),
					htmx.Get(paths.Backoffice()),
					htmx.Target("#modal-body"),
				),
			),
			Group(children),
		),
	)
}
