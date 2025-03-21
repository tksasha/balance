package component

import (
	"github.com/tksasha/balance/internal/app/balance"
	. "maragu.dev/gomponents"            //nolint:stylecheck
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *Component) Balance(balance *balance.Balance) Node {
	classes := Classes{
		"sum":          true,
		"text-danger":  balance.Balance < 0,
		"text-success": balance.Balance > 0,
		"fw-medium":    balance.Balance != 0,
	}

	return Div(Class("col-3"), ID("balance"),
		Div(Class("card cash"),
			Div(Class("card-body"),
				Table(Class("w-100"),
					TBody(
						Tr(Class("at_end"),
							Td(
								Text("Залишок на кінець"),
							),
							Td(Class("sum"),
								Text(c.Money(balance.AtEnd)),
							),
						),
						Tr(Class("balance"),
							Td(
								Text("Баланс"),
							),
							Td(classes,
								Text(c.Money(balance.Balance)),
							),
						),
					),
				),
			),
		),
	)
}
