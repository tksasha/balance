package models

import (
	"database/sql"
	"strconv"

	"github.com/tksasha/balance/pkg/currencies"
)

type Category struct {
	ID            int
	Name          string
	Income        bool
	Visible       bool
	Currency      currencies.Currency
	Supercategory int
	DeletedAt     sql.NullTime
}

type Categories []*Category

func (c *Category) GetIDAsString() string {
	return strconv.Itoa(c.ID)
}

func (c Categories) Income() Categories {
	categories := Categories{}

	for _, category := range c {
		if category.Income {
			categories = append(categories, category)
		}
	}

	return categories
}

func (c Categories) Expense() Categories {
	categories := Categories{}

	for _, category := range c {
		if !category.Income {
			categories = append(categories, category)
		}
	}

	return categories
}
