package component

import (
	"net/url"
	"strconv"
	"time"

	. "maragu.dev/gomponents" //nolint:stylecheck
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html" //nolint:stylecheck
)

const yearFrom = 2015

func (c *IndexComponent) Years(values url.Values) Node {
	var years []int

	for year := yearFrom; year <= time.Now().Year(); year++ {
		years = append(years, year)
	}

	return Div(
		ID("years"),
		htmx.SwapOOB("true"),
		Map(years, func(year int) Node {
			return c.year(year, values)
		}),
	)
}

func (c *IndexComponent) year(year int, values url.Values) Node {
	val, err := strconv.Atoi(values.Get("year"))
	current := err == nil && year == val

	classes := components.Classes{
		"active": current,
	}

	return A(
		classes,
		Href(c.ListItems(year, 0, values)),
		Text(strconv.Itoa(year)),
		htmx.Get(c.ListItems(year, 0, values)),
		htmx.Target("#items"),
	)
}
