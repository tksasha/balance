package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *Component) Index(cashes cash.Cashes, categories category.Categories, values url.Values) Node {
	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				Link(Rel("stylesheet"), Href("/assets/application.css")),
				Link(Rel("icon"), Type("image/x-icon"), Href("/assets/hryvnia.png")),
			},
			Body: []Node{
				c.header(values),
				If(true, c.cashes(cashes)),
				If(false, c.form(&item.Item{}, categories)),
				If(false, c.items()),
				If(
					true,
					Div(ID("items"), Style("display: none"),
						htmx.Get("/items/?month=02&year=2025"), // TODO: fix me
						htmx.Trigger("load"),
					),
				), // TODO: delme
				c.Modal(),
				Script(Src("/assets/bootstrap.js")),
				Script(Src("/assets/htmx.js")),
				Script(Src("/assets/application.js")),
			},
		},
	)
}
