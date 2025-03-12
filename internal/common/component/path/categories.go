package path

import (
	"net/url"
)

const categoriesPath = "/categories"

func Categories(values url.Values, params Params) string {
	values = setDefault(values)

	updateValues(values, params)

	path := url.URL{
		Path:     categoriesPath,
		RawQuery: values.Encode(),
	}

	return path.String()
}
