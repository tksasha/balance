package helpers

import (
	"net/url"
	"strconv"
)

func EditItemPath(id int) string {
	items := url.URL{
		Path: items,
	}

	items = *items.JoinPath(strconv.Itoa(id), "edit")

	return items.String()
}
