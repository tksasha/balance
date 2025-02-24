package component

import (
	"fmt"
	"maps"
	"net/url"
	"strconv"
	"time"
)

func (c *Component) ListItems(year, month int, values url.Values) string {
	path := url.URL{
		Path: "/items",
	}

	params := url.Values{}

	maps.Copy(params, values)

	params.Set("month", c.month(month, params))
	params.Set("year", c.year(year, params))

	path.RawQuery = params.Encode()

	return path.String()
}

func (c *Component) month(value int, values url.Values) string {
	var err error

	month := value

	if month == 0 {
		month, err = strconv.Atoi(values.Get("month"))
		if err != nil {
			month = int(time.Now().Month())
		}
	}

	return fmt.Sprintf("%02d", month)
}

func (c *Component) year(value int, values url.Values) string {
	var err error

	year := value

	if year == 0 {
		year, err = strconv.Atoi(values.Get("year"))
		if err != nil {
			year = time.Now().Year()
		}
	}

	return fmt.Sprintf("%04d", year)
}
