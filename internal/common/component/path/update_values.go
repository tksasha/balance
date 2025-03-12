package path

import (
	"net/url"
	"strconv"
	"time"
)

func updateValues(values url.Values, params Params) {
	if currency, ok := params["currency"]; ok {
		values.Set("currency", currency)
	}

	if month, ok := params["month"]; ok {
		values.Set("month", month)
	}

	if year, ok := params["year"]; ok {
		values.Set("year", year)
	}

	if _, ok := values["month"]; !ok {
		values.Set("month", strconv.Itoa(int(time.Now().Month())))
	}

	if _, ok := values["year"]; !ok {
		values.Set("year", strconv.Itoa(time.Now().Year()))
	}
}
