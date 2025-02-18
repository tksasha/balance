package helpers

import (
	"net/http"
	"net/url"

	"github.com/tksasha/balance/internal/core/common/valueobjects"
)

const items = "/items"

type Helpers struct {
	currentDateProvider valueobjects.CurrentDateProvider
}

func New(currentDateProvider valueobjects.CurrentDateProvider) *Helpers {
	return &Helpers{
		currentDateProvider: currentDateProvider,
	}
}

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
