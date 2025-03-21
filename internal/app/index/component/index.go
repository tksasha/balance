package component

import (
	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents"            //nolint:stylecheck
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

func (c *Component) Index(params params.Params) Node {
	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				Link(
					Rel("stylesheet"),
					Href("/assets/bootstrap-datepicker.min-cc8d75acca2a852e945170726a1ed804b63853ad.css"),
				),
				Link(Rel("stylesheet"), Href("/assets/application-9b924437db8302708b7ea7926e7b92072fbaec11.css")),
				Link(
					Rel("icon"),
					Type("image/x-icon"),
					Href("/assets/hryvnia-c8a5df41852f327f0ad50d3dddf29acf85597d22.png"),
				),
				Link(
					Rel("apple-touch-icon"),
					Href("/assets/hryvnia-c8a5df41852f327f0ad50d3dddf29acf85597d22.png"),
				),
			},
			Body: []Node{
				c.header(params),
				c.form(params),
				c.balance(params),
				c.categories(params),
				c.items(params),
				c.Modal(),
				c.linkToBackoffice(),
				Script(Src("/assets/bootstrap-0f43271223c74d330702ce94a39ed70d04e8fd36.js")),
				Script(Src("/assets/htmx.min-13dcd355b9ee9b169ddc7afea6683877be30920c.js")),
				Script(Src("/assets/jquery.min-5a9dcfbef655a2668e78baebeaa8dc6f41d8dabb.js")),
				Script(Src("/assets/bootstrap-datepicker.min-cdff2c53b8ff6b44eb16e842bd4b86541a7853f6.js")),
				Script(Src("/assets/bootstrap-datepicker.uk.min-d58d82ad3cc17da5ff61d0d8559c3b397c941638.js")),
				Script(Src("/assets/application-89a5eb026fbed807fc657de0661eafb71720a399.js")),
			},
		},
	)
}
