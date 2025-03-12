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

func EditCashPath(values url.Values, id int) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id), "edit").String()
}

func UpdateCashPath(values url.Values, id int) string {
	path := url.URL{
		Path:     cashesPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}
