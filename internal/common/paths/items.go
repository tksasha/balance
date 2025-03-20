package paths

import (
	"net/url"
	"strconv"
	"time"

	"github.com/tksasha/balance/internal/common/paths/params"
)

const itemsPath = "/items"

func Items(params params.Params) string {
	if !params.Has("month") {
		month := int(time.Now().Month())

		params.Set("month", strconv.Itoa(month))
	}

	if !params.Has("year") {
		params.Set("year", strconv.Itoa(time.Now().Year()))
	}

	path := url.URL{
		Path:     itemsPath,
		RawQuery: params.String(),
	}

	return path.String()
}

func UpdateItem(params params.Params, id int) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: params.String(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}

func NewItem(params params.Params) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: params.String(),
	}

	return path.JoinPath("new").String()
}

func CreateItem(params params.Params) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: params.String(),
	}

	return path.String()
}

func DeleteItem(params params.Params, id int) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: params.String(),
	}

	return path.JoinPath(strconv.Itoa(id)).String()
}

func EditItem(params params.Params, id int) string {
	path := url.URL{
		Path:     itemsPath,
		RawQuery: params.String(),
	}

	return path.JoinPath(strconv.Itoa(id), "edit").String()
}
