package path

import (
	"net/url"
	"strconv"
)

const backofficePath = "/backoffice"

func Backoffice() string {
	return "/backoffice"
}

func BackofficeCashes(params Params) string {
	values := setDefault(nil)

	updateCurrency(values, params)

	path := url.URL{
		Path:     backofficePath,
		RawQuery: values.Encode(),
	}

	return path.JoinPath(cashesPath).String()
}

func CreateBackofficeCash(params Params) string {
	return BackofficeCashes(params)
}

func BackofficeCash(params Params, id int) string {
	values := setDefault(nil)

	updateCurrency(values, params)

	path := url.URL{
		Path: backofficePath,
	}

	path = *path.JoinPath(cashesPath, strconv.Itoa(id))

	path.RawQuery = values.Encode()

	return path.String()
}

func UpdateBackofficeCash(params Params, id int) string {
	return BackofficeCash(params, id)
}

func DeleteBackofficeCash(params Params, id int) string {
	return BackofficeCash(params, id)
}

func EditBackofficeCash(params Params, id int) string {
	values := setDefault(nil)

	updateCurrency(values, params)

	path := url.URL{
		Path: backofficePath,
	}

	path = *path.JoinPath(cashesPath, strconv.Itoa(id), "edit")

	path.RawQuery = values.Encode()

	return path.String()
}

func BackofficeCategories() string {
	return Backoffice() + "/categories"
}
