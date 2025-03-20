package component

import (
	"fmt"
	"strconv"

	"github.com/tksasha/balance/internal/common/paths"
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
	}

	params.SetYear(year)

	callback := fmt.Sprintf("htmx.trigger('body', 'balance.year.changed', {balanceCategoriesPath: '%s'})",
		paths.Categories(params))

	value := strconv.Itoa(year)

	return Div(
		classes,
		Text(value),
		Data("number", value),
		htmx.Get(paths.Items(params)),
		htmx.Target("#items"),
		htmx.On(":after-request", callback),
	)
}
