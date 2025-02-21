package components

import (
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"      //nolint: stylecheck
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *ItemsComponent) List(items item.Items, months Node, years Node) Node {
	return Table(
		Class("table"),
		THead(
			Tr(
				Th(Class("text-center"), Text("Date")),
				Th(Class("text-center"), Text("Sum")),
				Th(Class("text-center"), Text("Category")),
				Th(Class("text-center"), Text("Description")),
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
			Class("text-center"),
			Text(c.Date(item.Date)),
		),
		Td(
			Class("text-end"),
			A(
				Href(c.EditItem(item.ID)),
				Text(c.Money(item.Sum)),
			),
		),
		Td(
			Text(item.CategoryName),
		),
		Td(
			Text(item.Description),
		),
	)
}
