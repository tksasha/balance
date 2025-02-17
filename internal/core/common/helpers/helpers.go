package helpers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const items = "/items"

func ItemsPath(r *http.Request, year, month int) string {
	u := base(items, r)

	values := u.Query()

	if year != 0 {
		values.Set("year", fmt.Sprintf("%04d", year))
	}

	if month != 0 {
		values.Set("month", fmt.Sprintf("%02d", month))
	}

	u.RawQuery = values.Encode()

	return u.String()
}

func EditItemPath(id int) string {
	items := url.URL{
		Path: items,
	}

	items = *items.JoinPath(strconv.Itoa(id), "edit")

	return items.String()
}

func base(path string, r *http.Request) url.URL {
	url := url.URL{
		Path: path,
	}

	if r == nil {
		return url
	}

	url.RawQuery = r.URL.Query().Encode()

	return url
}
