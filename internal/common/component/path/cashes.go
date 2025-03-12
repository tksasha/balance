package path

import (
	"net/url"
)

func Cashes(values url.Values) string {
	path := url.URL{
		Path:     "/cashes",
		RawQuery: setDefault(values).Encode(),
	}

	return path.String()
}
