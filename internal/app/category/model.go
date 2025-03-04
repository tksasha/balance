package category

import (
	"github.com/tksasha/balance/internal/common/currency"
)

type Category struct {
	ID            int
	Currency      currency.Currency
	Name          string
	Slug          string
	Income        bool
	Supercategory int
}

type Categories []*Category

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
