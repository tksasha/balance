package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) New(params path.Params, cash *cash.Cash) Node {
	return Div(
		c.Breadcrumbs(
			Li(Class("breadcrumb-item"),
				Span(Class("link"),
					htmx.Get(path.BackofficeCashes(params)),
					htmx.Target("#modal-body"),
					Text("Залишки"),
				),
			),
			Li(Class("breadcrumb-item active"), Text("Додавання")),
		),
		c.form(cash, nil),
	)
}
