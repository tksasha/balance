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
				Link(
					Rel("stylesheet"),
					Href("/assets/bootstrap-datepicker.min-cc8d75acca2a852e945170726a1ed804b63853ad.css"),
				),
				Link(
					Rel("icon"),
					Type("image/x-icon"),
					Href("/assets/hryvnia-c8a5df41852f327f0ad50d3dddf29acf85597d22.png"),
				),
				Link(Rel("stylesheet"), Href("/assets/application-9b924437db8302708b7ea7926e7b92072fbaec11.css")),
			},
			Body: []Node{
				If(true, c.header(values)),
				If(true, c.form()),
				If(true, c.balance()),
				If(true, c.categories()),
				If(true, c.items()),
				c.Modal(),
				Script(Src("/assets/bootstrap-0f43271223c74d330702ce94a39ed70d04e8fd36.js")),
				Script(Src("/assets/htmx.min-03a1ffdf83a11fab58acf6bcdf51233fdf14abd5.js")),
				Script(Src("/assets/jquery.min-5a9dcfbef655a2668e78baebeaa8dc6f41d8dabb.js")),
				Script(Src("/assets/bootstrap-datepicker.min-cdff2c53b8ff6b44eb16e842bd4b86541a7853f6.js")),
				Script(Src("/assets/bootstrap-datepicker.uk.min-d58d82ad3cc17da5ff61d0d8559c3b397c941638.js")),
				Script(Src("/assets/application-6f34a3d1859f17aa81029c6d93e9c48dfc999a12.js")),
			},
		},
	)
}
