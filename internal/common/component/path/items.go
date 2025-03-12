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

func EditItem(id int) string {
	return itemsPath + "/" + strconv.Itoa(id) + "/edit"
}

func UpdateItem(id int) string {
	return itemsPath + "/" + strconv.Itoa(id)
}

func NewItem() string {
	return itemsPath + "/new"
}

func CreateItem() string {
	return itemsPath
}

func DeleteItem(id int) string {
	return itemsPath + "/" + strconv.Itoa(id)
}
