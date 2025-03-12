package path

import (
	"net/url"
	"strconv"
)

const cashesPath = "/cashes"

func Cashes(values url.Values) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.String()
}

func EditCash(values url.Values, id int) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id), "edit").String()
}

func UpdateCash(values url.Values, id int) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}
