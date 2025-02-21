package components

import (
	"net/url"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents"            //nolint:stylecheck
	hx "maragu.dev/gomponents-htmx"      //nolint:stylecheck
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *IndexPageComponent) Index(categories category.Categories, values url.Values) Node {
	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				Link(Rel("stylesheet"), Href("/assets/bootstrap.min.css")),
				Link(Rel("stylesheet"), Href("/assets/application.css")),
			},
			Body: []Node{
				Div(
					Class("container mt-4 mb-4"),
					Div(
						Class("card mb-3"),
						Div(
							Class("card-body"),
							c.monthsComponent.Months(values),
						),
					),
					Div(
						Class("card mb-3"),
						Div(
							Class("card-body"),
							c.form(&item.Item{}, categories, nil),
						),
					),
					Div(
						Class("card"),
						Div(
							Class("card-body"),
							ID("items"),
							hx.Get("/items"),
							hx.Trigger("load"),
							Div(Class("spinner-border htmx-indicator"), ID("htmx-indicator")),
						),
					),
				),
				Script(Src("/assets/bootstrap.bundle.min.js")),
				Script(Src("/assets/htmx.min.js")),
			},
		},
	)
}
