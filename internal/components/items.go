package components

import (
	"github.com/tksasha/balance/internal/models"
	"maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func Items(items models.Items) gomponents.Node {
	return html.Table(
		html.Class("table"),
		html.THead(
			html.Tr(
				html.Th(html.Class("text-center"), gomponents.Text("Date")),
				html.Th(html.Class("text-center"), gomponents.Text("Sum")),
				html.Th(html.Class("text-center"), gomponents.Text("Category")),
				html.Th(html.Class("text-center"), gomponents.Text("Description")),
			),
		),
		html.TBody(
			gomponents.Map(
				items,
				func(item *models.Item) gomponents.Node {
					return html.Tr(
						html.Td(html.Class("text-center"), gomponents.Text(item.GetDateAsString())),
						html.Td(
							html.Class("text-end"),
							html.A(
								html.Href("/items"),
								hx.Get("/items"),
								gomponents.Text(item.GetSumAsString()),
							),
						),
						html.Td(gomponents.Text(item.CategoryName)),
						html.Td(gomponents.Text(item.Description)),
					)
				},
			),
		),
	)
}
