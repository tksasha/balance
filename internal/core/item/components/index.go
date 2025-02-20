package components

import (
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents" //nolint: stylecheck
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint: stylecheck
)

func (c *ItemsComponent) Index(items item.Items) Node {
	return Span(
		Table(
			Class("table"),
			ID("items"),
			hx.SwapOOB("true"),
			THead(
				Tr(
					Th(Class("text-center"), Text("Date")),
					Th(Class("text-center"), Text("Sum")),
					Th(Class("text-center"), Text("Category")),
					Th(Class("text-center"), Text("Description")),
				),
			),
			TBody(
				Map(items, c.row),
			),
		),
		Text("42"),
	)
}

func (c *ItemsComponent) row(item *item.Item) Node {
	return Tr(
		Td(
			Class("text-center"),
			Text(c.Date(item.Date)),
		),
		Td(
			Class("text-end"),
			A(
				Href(c.Helpers.EditItemPath(item.ID)),
				Text(sum(item.Sum)),
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
