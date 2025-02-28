package component

import (
	"net/url"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/category"
	"github.com/tksasha/balance/internal/app/item"
	. "maragu.dev/gomponents"            //nolint:stylecheck
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
				If(false, c.header(values)),
				If(true, c.cashComponent.List(cashes)),
				If(false, c.form(&item.Item{}, categories)),
				If(false, c.items()),
				c.Modal(),
				Script(Src("/assets/bootstrap.js")),
				Script(Src("/assets/htmx.js")),
				Script(Src("/assets/application.js")),
			},
		},
	)
}
