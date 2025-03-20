package paths

import (
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/common/paths/params"
)

const backofficePath = "/backoffice"

func Backoffice() string {
	return "/backoffice"
}

func BackofficeCashes(params params.Params) string {
	path := url.URL{
		Path:     backofficePath,
		RawQuery: params.String(),
	}

	return path.JoinPath(cashesPath).String()
}

func EditBackofficeCash(id int) string {
	path := url.URL{
		Path: backofficePath,
	}

	return path.JoinPath(cashesPath, strconv.Itoa(id), "edit").String()
}

func BackofficeCash(id int) string {
	path := url.URL{
		Path: backofficePath,
	}

	return path.JoinPath(cashesPath, strconv.Itoa(id)).String()
}

func DeleteBackofficeCash(id int) string {
	return BackofficeCash(id)
}

func UpdateBackofficeCash(id int) string {
	return BackofficeCash(id)
}

func NewBackofficeCash(params params.Params) string {
	path := url.URL{
		Path:     backofficePath,
		RawQuery: params.String(),
	}

	return path.JoinPath(cashesPath, "new").String()
}

func CreateBackofficeCash() string {
	path := url.URL{
		Path: backofficePath,
	}

	return path.JoinPath(cashesPath).String()
}

func BackofficeCategories(params params.Params) string {
	path := url.URL{
		Path:     backofficePath,
		RawQuery: params.String(),
	}

	return path.JoinPath(categoriesPath).String()
}

func NewBackofficeCategory(params params.Params) string {
	path := url.URL{
		Path:     backofficePath,
		RawQuery: params.String(),
	}

	return path.JoinPath(categoriesPath, "new").String()
}

func CreateBackofficeCategory() string {
	path := url.URL{
		Path: backofficePath,
	}

	return path.JoinPath(categoriesPath).String()
}

func EditBackofficeCategory(id int) string {
	path := url.URL{
		Path: backofficePath,
	}

	return path.JoinPath(categoriesPath, strconv.Itoa(id), "edit").String()
}

func UpdateBackofficeCategory(id int) string {
	path := url.URL{
		Path: backofficePath,
	}

	return path.JoinPath(categoriesPath, strconv.Itoa(id)).String()
}
