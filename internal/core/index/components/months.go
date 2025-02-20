package components

import (
	"fmt"
	"net/http"

	"github.com/tksasha/balance/internal/core/common/helpers"
	"github.com/tksasha/month"
	. "maragu.dev/gomponents" //nolint:stylecheck
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type MonthsComponent struct {
	helpers *helpers.Helpers
}

func NewMonthsComponent(helpers *helpers.Helpers) *MonthsComponent {
	return &MonthsComponent{
		helpers: helpers,
	}
}

func (c *MonthsComponent) Months(req *http.Request) Node {
	row := func(month month.Month) Node {
		return c.Month(req, month)
	}

	return Div(
		Map(month.All(), row),
	)
}

func (c *MonthsComponent) Month(req *http.Request, month month.Month) Node {
	return A(
		If(
			c.isActive(req.URL.Query().Get("month"), month.Number),
			Class("active"),
		),
		Href(c.helpers.ItemsPath(req, 0, month.Number)),
		Text(month.Name),
		hx.Get(c.helpers.ItemsPath(req, 0, month.Number)),
	)
}

func (c *MonthsComponent) isActive(active string, number int) bool {
	return active == fmt.Sprintf("%02d", number)
}
