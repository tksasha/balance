package components

import (
	"github.com/tksasha/balance/internal/models"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Items(items models.Items) gomponents.Node {
	return html.Table(
		html.Class("table"),
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("ID")),
				html.Th(gomponents.Text("Date")),
				html.Th(gomponents.Text("Sum")),
				html.Th(gomponents.Text("Category")),
				html.Th(gomponents.Text("Description")),
			),
		),
		html.TBody(
			gomponents.Map(
				items,
				func(item *models.Item) gomponents.Node {
					return html.Tr(
						html.Td(gomponents.Text(item.GetIDAsString())),
						html.Td(gomponents.Text(item.GetDateAsString())),
						html.Td(gomponents.Text(item.GetSumAsString())),
						html.Td(gomponents.Text(item.CategoryName)),
						html.Td(gomponents.Text(item.Description)),
					)
				},
			),
		),
	)
}
