package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) cash(cash *cash.Cash) Node {
	return Tr(ID(c.cashID(cash.ID)), htmx.SwapOOB("true"),
		Td(Text(cash.Name)),
		Td(Class("sum"),
			Div(Class("link"),
				htmx.Get(c.cashEditPath(cash.ID)),
				htmx.Target("#modal-body"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
				Text(c.Money(cash.Sum)),
			),
		),
	)
}
