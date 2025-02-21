package components

import (
	"net/url"

	"github.com/tksasha/balance/internal/core/category"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *IndexComponent) Index(categories category.Categories, values url.Values) Node {
	_ = c.form

	css := "/assets/application-28267bd7a3abb90cde33ecb57c0adb229888a59258f97d91a688ed1377a381c1.css"

	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				// Link(Rel("stylesheet"), Href("/assets/bootstrap.min.css")),
				Link(Rel("stylesheet"), Href(css)),
				// Link(Rel("stylesheet"), Href("/assets/application.css")),
			},
			Body: []Node{
				Header(
					c.monthsComponent.Months(values),
					c.yearsComponent.Years(values),
				),
				// Div(
				// 	Class("card mb-3"),
				// 	Div(
				// 		Class("card-body"),
				// 		c.form(&item.Item{}, categories, nil),
				// 	),
				// ),
				Div(
					Class("container"),
					Div(
						Class("row mt-4 mb-5"),
						Div(
							Class("col"),
							Div(
								Class("card items"),
								Div(
									Class("card-body"),
									ID("items"),
									htmx.Get(c.ListItems(0, 0, url.Values{})),
									htmx.Trigger("load"),
									Div(Class("spinner-border htmx-indicator"), ID("htmx-indicator")),
								),
							),
						),
					),
				),
				Script(Src("/assets/bootstrap.bundle.min.js")),
				Script(Src("/assets/htmx.min.js")),
			},
		},
	)
}
