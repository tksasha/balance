package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Edit(params params.Params, cash *cash.Cash) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(paths.BackofficeCashes(params)),
					htmx.Target("#modal-body"),
					Text("Залишки"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Редагування")),
		),
		c.form(cash, nil),
	)
}
