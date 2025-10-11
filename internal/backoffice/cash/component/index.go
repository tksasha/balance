package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/paths"
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:staticcheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:staticcheck
)

func (c *Component) Index(parameters params.Params, cashes cash.Cashes) Node {
	return Div(
		Div(Class("d-flex justify-content-between"),
			c.Breadcrumbs(Li(Class("breadcrumb-item active"), Text("Залишки"))),
			Div(
				Button(Class("btn btn-outline-primary btn-sm"),
					Text("Додати"),
					htmx.Get(paths.NewBackofficeCash(parameters)),
					htmx.Target("#modal-body"),
				),
			),
		),
		Form(Class("mb-3"),
			Select(Class("form-select form-select-sm"), Name("currency"),
				htmx.Get(paths.BackofficeCashes(params.Params{})),
				htmx.Target("#modal-body"),
				c.CurrencyOptions(parameters.Get("currency")),
			),
		),
		Table(Class("table table-borderless"),
			THead(
				Tr(
					Th(Text("Назва")),
					Th(Class("text-end"), Text("Сума")),
					Th(Class("text-end"), Text("Група")),
				),
			),
			TBody(Div(Map(cashes, c.cash))),
		),
	)
}

func (c *Component) cash(cash *cash.Cash) Node {
	return Tr(
		Td(Div(Class("link"),
			htmx.Get(paths.EditBackofficeCash(cash.ID)),
			htmx.Target("#modal-body"),
			Text(cash.Name)),
		),
		Td(Class("text-end"), Text(c.Money(cash.Sum))),
		Td(Class("text-end text-secondary"), c.Text(cash.Supercategory)),
	)
}
