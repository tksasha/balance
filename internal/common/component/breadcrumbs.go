package component

import (
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Breadcrumbs(children ...Node) Node {
	return Nav(Class("mb-4"),
		Ol(Class("breadcrumb"),
			Li(Class("breadcrumb-item"),
				Span(Class("link"), Text("Backoffice"),
					htmx.Get(path.Backoffice()),
					htmx.Target("#modal-body"),
				),
			),
			Group(children),
		),
	)
}
