package component

import (
	"github.com/tksasha/balance/internal/app/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) cashes(residual, balance float64, cashes cash.Cashes) Node {
	return Div(Class("container-fluid"),
		Div(Class("row mt-4"),
			c.balanceComponent.Balance(residual, balance),
			c.cashComponent.List(cashes),
		),
	)
}
