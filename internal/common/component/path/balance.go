package path

import (
	"net/url"
)

func Balance(values url.Values) string {
	path := url.URL{
		Path:     "/balance",
		RawQuery: setDefault(values).Encode(),
	}

	return path.String()
}
