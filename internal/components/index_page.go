package components

//nolint:stylecheck // ST1001
import (
	"github.com/tksasha/balance/internal/models"
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func IndexPage() Node {
	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				Stylesheet("/assets/bootstrap.min.css"),
				Stylesheet("/assets/application.css"),
			},
			Body: []Node{
				Div(
					Class("container mt-4"),
					Div(
						Class("card mb-3"),
						Div(
							Class("card-body"),
							ItemForm(&models.Item{}, nil),
						),
					),
					Div(
						Class("card"),
						Div(
							Class("card-body"),
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
