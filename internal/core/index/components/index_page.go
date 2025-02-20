package components

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/common/components"
	"github.com/tksasha/balance/internal/core/item"
	. "maragu.dev/gomponents" //nolint:stylecheck
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/components" //nolint:stylecheck
	. "maragu.dev/gomponents/html"       //nolint:stylecheck
)

type IndexPageComponent struct {
	components.BaseComponent

	monthsComponent *MonthsComponent
}

func NewIndexPageComponent(monthsComponent *MonthsComponent) *IndexPageComponent {
	return &IndexPageComponent{
		monthsComponent: monthsComponent,
	}
}

func (c *IndexPageComponent) Index(req *http.Request, categories category.Categories) Node {
	return HTML5(
		HTML5Props{
			Title:    "Balance",
			Language: "en",
			Head: []Node{
				components.Stylesheet("/assets/bootstrap.min.css"),
				components.Stylesheet("/assets/application.css"),
			},
			Body: []Node{
				Div(
					Class("container mt-4 mb-4"),
					Div(
						Class("card mb-3"),
						Div(
							Class("card-body"),
							c.monthsComponent.Months(req),
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
