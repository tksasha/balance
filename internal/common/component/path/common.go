package path

import (
	"maps"
	"net/url"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/common/currency"
)

func setDefault(values url.Values) url.Values {
	params := make(url.Values, len(values))

	maps.Copy(params, values)

	if !params.Has("currency") {
		params.Add("currency", currency.GetCode(currency.Default))
	}

	return params
}

func updateCurrency(values url.Values, params Params) {
	if currency, ok := params["currency"]; ok {
		values.Set("currency", currency)
	}
}

func updateValues(values url.Values, params Params) {
	updateCurrency(values, params)

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
