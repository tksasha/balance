package component

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (c *Component) ListItems(year, month int, req *http.Request) string {
	path := url.URL{
		Path: "/items",
	}

	values := path.Query()

	values.Set("month", c.month(month, req))
	values.Set("year", c.year(year, req))

	path.RawQuery = values.Encode()

	return path.String()
}

func (c *Component) month(value int, req *http.Request) string {
	var err error

	month := value

	if month == 0 {
		month, err = strconv.Atoi(req.URL.Query().Get("month"))
		if err != nil {
			month = int(time.Now().Month())
		}
	}

	return fmt.Sprintf("%02d", month)
}

func (c *Component) year(value int, req *http.Request) string {
	var err error

	year := value

	if year == 0 {
		year, err = strconv.Atoi(req.URL.Query().Get("year"))
		if err != nil {
			year = time.Now().Year()
		}
	}

	return fmt.Sprintf("%04d", year)
}
