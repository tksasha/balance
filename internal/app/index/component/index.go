package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/category"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *IndexComponent) Index(categories category.Categories, values url.Values) Node {
	_ = c.form

	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				Link(Rel("stylesheet"), Href("/assets/application.css")),
				Link(Rel("icon"), Type("image/x-icon"), Href("/assets/hryvnia.png")),
			},
			Body: []Node{
				Header(
					c.Months(values),
					c.Years(values),
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
				c.Modal(),
				Script(Src("/assets/bootstrap.js")),
				Script(Src("/assets/htmx.js")),
			},
		},
	)
}
