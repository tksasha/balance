package components

import (
	"fmt"
	"net/url"

	"github.com/tksasha/balance/internal/core/common/component"
	"github.com/tksasha/month"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type MonthsComponent struct {
	*component.Component
}

func NewMonthsComponent(component *component.Component) *MonthsComponent {
	return &MonthsComponent{
		Component: component,
	}
}

func (c *MonthsComponent) Months(values url.Values) Node {
	return Div(
		ID("months"),
		htmx.SwapOOB("true"),
		Map(month.All(), func(month month.Month) Node {
			return c.Month(month, values)
		}),
	)
}

func (c *MonthsComponent) Month(month month.Month, values url.Values) Node {
	return A(
		If(
			c.isActive(values.Get("month"), month.Number),
			Class("active"),
		),
		Href(c.ListItems(0, month.Number, values)),
		Text(month.Name),
		htmx.Get(c.ListItems(0, month.Number, values)),
		htmx.Target("#items"),
	)
}

func (c *MonthsComponent) isActive(active string, number int) bool {
	return active == fmt.Sprintf("%02d", number)
}
