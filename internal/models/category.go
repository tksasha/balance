package models

type Category struct {
	ID     int
	Name   string
	Income bool
}

type Categories []*Category
