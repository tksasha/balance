package path

import (
	"net/url"
	"strconv"
	"time"
)

const categoriesPath = "/categories"

func Categories(params Params) string {
	path := url.URL{Path: categoriesPath}

	values := url.Values{}

	month, ok := params["month"]
	if ok {
		values.Add("month", month)
	} else {
		values.Add("month", strconv.Itoa(int(time.Now().Month())))
	}

	year, ok := params["year"]
	if ok {
		values.Add("year", year)
	} else {
		values.Add("year", strconv.Itoa(time.Now().Year()))
	}

	path.RawQuery = values.Encode()

	return path.String()
}
