package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/backoffice/cash"
	. "maragu.dev/gomponents"      //nolint:stylecheck
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	return Div(
		c.Breadcrumbs(Li(Class("breadcrumb-item active"), Text("Залишки"))),
		Table(Class("w-100"),
			THead(
				Tr(
					Th(Text("Назва")),
					Th(Text("Сума")),
					Th(Text("Категорія")),
					Th(Text("Обраний")),
					Th(),
				),
			),
			TBody(
				Map(cashes, c.cash),
			),
		),
	)
}

func (c *Component) cash(cash *cash.Cash) Node {
	return Tr(
		Td(Class("pe-3 pb-2"),
			Input(Class("form-control"), Value(cash.Name)),
		),
		Td(Class("pe-3"),
			Input(Class("form-control text-end"), Value(cash.Formula)),
		),
		Td(Class("pe-3"),
			Input(Class("form-control text-end"), Value(strconv.Itoa(cash.Supercategory))),
		),
		Td(Class("pe-3 text-center"),
			Input(Type("checkbox"), Class("form-check-input")),
		),
		Td(Class("text-end"),
			Button(Class("btn-close"), Type("button"), Title("Видалити")),
		),
	)
}
