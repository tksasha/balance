package component

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/common/component/path"
	"github.com/tksasha/month"
	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

func (c *Component) Months(values url.Values) Node {
	return Div(
		ID("months"),
		htmx.SwapOOB("true"),
		Map(month.All(), func(month month.Month) Node {
			return c.Month(month, values)
		}),
	)
}

func (c *Component) Month(month month.Month, values url.Values) Node {
	val, err := strconv.Atoi(values.Get("month"))
	current := err == nil && month.Number == val

	classes := components.Classes{
		"active": current,
		"link":   true,
	}

	number := strconv.Itoa(month.Number)

	params := path.Params{"month": number}

	callback := fmt.Sprintf("htmx.trigger('body', 'balance.month.changed', {balanceCategoriesPath: '%s'})",
		path.Categories(values, params))

	return Div(
		classes,
		Text(month.Name),
		Data("number", number),
		htmx.Get(path.Items(values, params)),
		htmx.Target("#items"),
		htmx.On(":after-request", callback),
	)
}
