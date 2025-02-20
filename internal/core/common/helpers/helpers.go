package helpers

import (
	"net/http"
	"net/url"
)

const items = "/items"

func base(path string, request *http.Request) url.URL {
	url := url.URL{
		Path: path,
	}

	if request == nil {
		return url
	}

	url.RawQuery = request.URL.Query().Encode()

	return url
}
