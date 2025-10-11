package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) New(params params.Params, cash *cash.Cash) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(paths.BackofficeCashes(params)),
					htmx.Target("#modal-body"),
					Text("Залишки"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Додавання")),
		),
		c.form(cash, nil),
	)
}
