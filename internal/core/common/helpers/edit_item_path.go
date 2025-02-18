package helpers

import (
	"net/url"
	"strconv"
)

func (h *Helpers) EditItemPath(id int) string {
	items := url.URL{
		Path: items,
	}

	items = *items.JoinPath(strconv.Itoa(id), "edit")

	return items.String()
}
