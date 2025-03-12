package path

import (
	"net/url"
	"strconv"
)

const itemsPath = "/items"

func Items(values url.Values, params Params) string {
	values = setDefault(values)

	updateValues(values, params)

	path := url.URL{
		Path:     itemsPath,
		RawQuery: values.Encode(),
	}

	return path.String()
}

func EditItem(values url.Values, id int) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id), "edit").String()
}

func UpdateItem(values url.Values, id int) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}

func NewItem(values url.Values) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath("new").String()
}

func CreateItem(values url.Values) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.String()
}

func DeleteItem(values url.Values, id int) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: setDefault(values).Encode(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}
