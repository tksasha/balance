package component

import (
	"net/url"

	. "maragu.dev/gomponents"            //nolint:stylecheck
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *Component) Index(values url.Values) Node {
	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				Link(Rel("stylesheet"), Href("/assets/application-a98def7283f33f69527a05b02b65a2f2c8b52d7d.css")),
				Link(Rel("icon"), Type("image/x-icon"), Href("/assets/hryvnia.png")),
			},
			Body: []Node{
				If(false, c.header(values)),
				If(true, c.balance()),
				If(true, c.categories()),
				If(true, c.form()),
				If(false, c.items()),
				c.Modal(),
				Script(Src("/assets/bootstrap-0f43271223c74d330702ce94a39ed70d04e8fd36.js")),
				Script(Src("/assets/htmx-ac810f4cc51114714079b5051f1bb57802a9625b.js")),
				Script(Src("/assets/application-02005eb566d0689befc6cffc9b68fa08e3524d8d.js")),
			},
		},
	)
}
