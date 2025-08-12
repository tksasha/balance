package component

import (
	"strconv"

	"github.com/tksasha/balance/internal/common/paths/params"
	"github.com/tksasha/month"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Months(params params.Params) Node {
	return Div(
		ID("months"),
		htmx.SwapOOB("true"),
		Map(month.All(), func(month month.Month) Node {
			return c.Month(month, params)
		}),
	)
}

func (c *Component) Month(month month.Month, params params.Params) Node {
	val, err := strconv.Atoi(params.Get("month"))
	current := err == nil && month.Number == val

	classes := components.Classes{
		"active": current,
		"link":   true,
		"today":  c.timeProvider.IsCurrentMonth(month.Number),
	}

	return Div(
		classes,
		Text(month.Name),
		Data("number", strconv.Itoa(month.Number)),
	)
}
