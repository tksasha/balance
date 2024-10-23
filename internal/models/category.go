package models

import "strconv"

type Category struct {
	ID     int
	Name   string
	Income bool
}

type Categories []*Category

func (c *Category) GetIDAsString() string {
	return strconv.Itoa(c.ID)
}
