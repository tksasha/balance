package component

import (
	"github.com/tksasha/balance/internal/backoffice/cash"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) List(cashes cash.Cashes) Node {
	return Div(
		Div(Class("d-flex justify-content-between"),
			c.Breadcrumbs(Li(Class("breadcrumb-item active"), Text("Залишки"))),
			Div(
				Button(Class("btn btn-outline-primary btn-sm"), Text("Додати")),
			),
		),
		Form(Class("mb-3"),
			Select(Class("form-select form-select-sm"), Name("currency"),
				htmx.Get(path.BackofficeCashes()),
				htmx.Target("#modal-body"),
				Option(Text("UAH"), Value("uah")),
				Option(Text("USD"), Value("usd")),
				Option(Text("EUR"), Value("eur")),
			),
		),
		Table(Class("table table-borderless"),
			TBody(Div(Map(cashes, c.cash))),
		),
	)
}

func (c *Component) cash(cash *cash.Cash) Node {
	return Tr(
		Td(Text(cash.Name)),
		Td(Class("text-end"),
			Div(Class("link"),
				htmx.Get(path.BackofficeEditCash(cash.ID)),
				htmx.Target("#modal-body"),
				Text(c.Money(cash.Sum))),
		),
	)
}
