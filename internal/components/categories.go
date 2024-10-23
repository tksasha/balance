package components

import (
	"github.com/tksasha/balance/internal/models"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Categories(categories models.Categories) gomponents.Node {
	return html.Table(
		html.Class("table"),
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("ID")),
				html.Th(gomponents.Text("Name")),
			),
		),
		html.TBody(
			gomponents.Map(
				categories,
				func(category *models.Category) gomponents.Node {
					return html.Tr(
						html.Td(gomponents.Text(category.GetIDAsString())),
						html.Td(gomponents.Text(category.Name)),
					)
				},
			),
		),
	)
}
