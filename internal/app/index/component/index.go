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
					Href("/assets/datepicker.min-1b98f51332bc364984c9fc3291cc6e20188fefb0.css")),
				Link(
					Rel("stylesheet"),
					Href("/assets/datepicker-bs5.min-f2c1cb83e2274f112ab4841abeb7303ea290a018.css")),
				Link(
					Rel("stylesheet"),
					Href("/assets/application-26eb1260657e8b37e8a5de9e77b593bc47fdd748.css")),
				Link(
					Rel("icon"),
					Type("image/x-icon"),
					Href("/assets/hryvnia-c8a5df41852f327f0ad50d3dddf29acf85597d22.png")),
				Link(
					Rel("apple-touch-icon"),
					Href("/assets/hryvnia-c8a5df41852f327f0ad50d3dddf29acf85597d22.png")),
			},
			Body: []Node{
				c.header(params),
				c.form(params),
				c.balance(params),
				c.categories(),
				c.items(),
				c.Modal(),
				c.linkToBackoffice(),
				c.currenciesWidget(params),
				c.currenciesWidgetContent(params),
				Script(Src("/assets/popper-67b2805cfefc6e46c7e47616477d64e048dcec63.js")),
				Script(Src("/assets/bootstrap-0f43271223c74d330702ce94a39ed70d04e8fd36.js")),
				Script(Src("/assets/htmx.min-13dcd355b9ee9b169ddc7afea6683877be30920c.js")),
				Script(Src("/assets/datepicker.min-c069e98468d2389c7735d150d4ec164f51033dd4.js")),
				Script(Src("/assets/application-bc59ff578d9937462689ce70543d30071c500301.js")),
				// Script(Src("/assets/application.js?" + strconv.Itoa(int(time.Now().Unix())))),
			},
		},
	)
}
