package component

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/common/component/path"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Years(values url.Values) Node {
	return Div(
		ID("years"),
		htmx.SwapOOB("true"),
		Map(c.years, func(year int) Node {
			return c.year(year, values)
		}),
	)
}

func (c *Component) year(year int, values url.Values) Node {
	val, err := strconv.Atoi(values.Get("year"))
	current := err == nil && year == val

	classes := components.Classes{
		"active": current,
		"link":   true,
	}

	number := strconv.Itoa(year)

	params := path.Params{"year": number}

	callback := fmt.Sprintf("htmx.trigger('body', 'balance.year.changed', {balanceCategoriesPath: '%s'})",
		path.Categories(values, params))

	return Div(
		classes,
		Text(strconv.Itoa(year)),
		Data("number", number),
		htmx.Get(path.Items(values, params)),
		htmx.Target("#items"),
		htmx.On(":after-request", callback),
	)
}
