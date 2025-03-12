package path

import (
	"maps"
	"net/url"

	"github.com/tksasha/balance/internal/common/currency"
)

func Balance(values url.Values) string {
	params := make(url.Values, len(values))

	maps.Copy(params, values)

	if !params.Has("currency") {
		params.Add("currency", currency.GetCode(currency.Default))
	}

	path := url.URL{
		Path:     "/balance",
		RawQuery: params.Encode(),
	}

	return path.String()
}
