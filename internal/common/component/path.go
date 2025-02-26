package component

import (
	"net/url"
	"strconv"
)

func (c *Component) EditCash(id int) string {
	path := &url.URL{
		Path: "/cashes",
	}

	path = path.JoinPath(strconv.Itoa(id), "edit")

	return path.String()
}
