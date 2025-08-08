package component

import (
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/common/paths/params"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Years(params params.Params) Node {
	return Div(
		ID("years"),
		htmx.SwapOOB("true"),
		Map(c.years, func(year int) Node {
			return c.year(year, params)
		}),
	)
}

func (c *Component) year(year int, params params.Params) Node {
	val, err := strconv.Atoi(params.Get("year"))
	current := err == nil && year == val

	classes := components.Classes{
		"active": current,
		"link":   true,
		"today":  time.Now().Year() == year,
	}

	value := strconv.Itoa(year)

	return Div(
		classes,
		Text(value),
		Data("number", value),
	)
}
