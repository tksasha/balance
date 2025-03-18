package component

import (
	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Index(params path.Params, categories category.Categories) Node {
	return Div(
		Div(Class("d-flex justify-content-between"),
			c.Breadcrumbs(Li(Class("breadcrumb-item active"), Text("Категорії"))),
			Div(
				Button(Class("btn btn-outline-primary btn-sm"),
					Text("Додати"),
					htmx.Get(path.NewBackofficeCategory(params)),
					htmx.Target("#modal-body"),
				),
			),
		),
		Form(Class("mb-3"),
			Select(Class("form-select form-select-sm"), Name("currency"),
				htmx.Get("/backoffice/categories"), // TODO: fix me
				htmx.Target("#modal-body"),
				c.CurrencyOptions(params["currency"]),
			),
		),
		Table(Class("table table-borderless table-hover"),
			THead(Tr(
				Th(Text("Назва")),
				Th(Class("text-center"), Text("Надходження")),
				Th(Class("text-center"), Text("Відображається")),
			)),
			TBody(Div(Map(categories, c.category))),
		),
	)
}

func (c *Component) category(category *category.Category) Node {
	return Tr(
		Td(
			Div(Class("link"), Text(category.Name)),
			htmx.Get(path.EditBackofficeCategory(path.NewCurrency(category.Currency), category.ID)),
			htmx.Target("#modal-body"),
		),
		Td(Class("text-center"),
			If(category.Income, c.Yes()),
			If(!category.Income, c.No()),
		),
		Td(Class("text-center"),
			If(category.Visible, c.Yes()),
			If(!category.Visible, c.No()),
		),
	)
}
