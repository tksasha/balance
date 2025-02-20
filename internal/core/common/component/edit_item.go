package component

import (
	"net/url"
	"strconv"
)

func (c *Component) EditItem(id int) string {
	path := &url.URL{
		Path: "/items",
	}

	path = path.JoinPath(strconv.Itoa(id), "edit")

	return path.String()
}
