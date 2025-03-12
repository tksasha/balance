package path

import (
	"maps"
	"net/url"

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
