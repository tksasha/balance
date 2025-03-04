package component

import (
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint: stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *Component) List(items item.Items, months Node, years Node) Node {
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
		TBody(Map(items, c.item)),
		Iff(months != nil, func() Node { return months }),
		Iff(years != nil, func() Node { return years }),
	)
}

func (c *Component) item(item *item.Item) Node {
	return c.Template(
		Tr(
			Td(Class("items-date"),
				Text(date(item.Date)),
			),
			Td(Class("items-sum"),
				Div(Class("text-primary"),
					Text(c.Money(item.Sum)),
					Style("cursor: pointer"),
					htmx.Get(c.editPath(item.ID)),
					htmx.Target("#modal-body"),
					htmx.Trigger("click"),
					Data("bs-toggle", "modal"),
					Data("bs-target", "#modal"),
				),
			),
			Td(Class("items-category"),
				Text(item.CategoryName),
			),
			Td(c.Description(item.Description)),
		),
	)
}
