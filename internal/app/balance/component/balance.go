package component

import (
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Balance(residual, balance float64) Node {
	return Div(Class("col-3"),
		Div(Class("card cash"),
			Div(Class("card-body"),
				Table(Class("w-100"),
					TBody(
						Tr(Class("at_end"),
							Td(
								Text("Залишок на кінець"),
							),
							Td(Class("sum"),
								Div(Class("link"),
									Text(c.Money(residual)),
								),
							),
						),
						Tr(Class("balance"),
							Td(
								Text("Баланс"),
							),
							Td(Class("sum"),
								Div(Class("link"),
									Text(c.Money(balance)),
								),
							),
						),
					),
				),
			),
		),
	)
}
