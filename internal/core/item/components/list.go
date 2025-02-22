package components

import (
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *ItemsComponent) List(items item.Items, months Node, years Node) Node {
	return Table(
		Class("table table-hover table-borderless"),
		THead(
			Tr(
				Th(Class("items-date"), Text("Дата")),
				Th(Class("items-sum"), Text("Сума")),
				Th(Class("items-category"), Text("Категорія")),
				Th(Text("Опис")),
			),
		),
		TBody(
			Map(items, c.item),
		),
		Iff(months != nil, func() Node { return months }),
		Iff(years != nil, func() Node { return years }),
	)
}

func (c *ItemsComponent) item(item *item.Item) Node {
	return Tr(
		Td(
			Class("items-date"),
			Text(c.Date(item.Date)),
		),
		Td(
			Class("items-sum"),
			Div(
				Class("text-primary"),
				Text(c.Money(item.Sum)),
				Style("cursor: pointer"),
				htmx.Get(c.EditItem(item.ID)),
				htmx.Target("#modal-body"),
				htmx.Trigger("click"),
				Data("bs-toggle", "modal"),
				Data("bs-target", "#modal"),
			),
		),
		Td(
			Class("items-category"),
			Text(item.CategoryName.String),
		),
		Td(
			Raw(c.Description(item.Description)),
		),
	)
}
