package components

import (
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/common/component"
	"github.com/tksasha/month"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

type MonthsComponent struct {
	*component.Component
}

func NewMonthsComponent() *MonthsComponent {
	return &MonthsComponent{
		Component: component.New(),
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
	val, err := strconv.Atoi(values.Get("month"))
	current := err == nil && month.Number == val

	classes := components.Classes{
		"active": current,
	}

	return A(
		classes,
		Href(c.ListItems(0, month.Number, values)),
		Text(month.Name),
		htmx.Get(c.ListItems(0, month.Number, values)),
		htmx.Target("#items"),
	)
}
