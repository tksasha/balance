package components

//nolint:stylecheck // ST1001
import (
	"github.com/tksasha/balance/internal/models"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func Items(items models.Items) Node {
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
			Map(
				items,
				func(item *models.Item) Node {
					return Tr(
						Td(Class("text-center"), Text(item.GetDateAsString())),
						Td(
							Class("text-end"),
							A(
								Href("/items"),
								hx.Get("/items"),
								Text(item.GetSumAsString()),
							),
						),
						Td(Text(item.CategoryName)),
						Td(Text(item.Description)),
					)
				},
			),
		),
	)
}
