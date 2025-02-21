package component

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

func (c *Component) ListItems(year, month int, values url.Values) string {
	path := url.URL{
		Path: "/items",
	}

	values.Set("month", c.month(month, values))
	values.Set("year", c.year(year, values))

	path.RawQuery = values.Encode()

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
