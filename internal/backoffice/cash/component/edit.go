package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Edit(cash *cash.Cash) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(path.BackofficeCashes()),
					htmx.Target("#modal-body"),
					Text("Залишки"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Редагування")),
		),
		c.form(cash, nil),
	)
}
