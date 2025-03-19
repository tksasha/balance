package category

import (
	"maps"
	"slices"

	"github.com/tksasha/balance/internal/common/currency"
)

type Category struct {
	ID            int
	Currency      currency.Currency
	Name          string
	Income        bool
	Supercategory int
	Sum           float64
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

func (c Categories) HasMoreThanOne() bool {
	return len(c) > 1
}

func (c Categories) Sum() float64 {
	var sum float64

	for _, entity := range c {
		sum += entity.Sum
	}

	return sum
}

type GroupedCategories map[int]Categories

func (e GroupedCategories) Keys() []int {
	keys := slices.Collect(maps.Keys(e))

	slices.Sort(keys)

	return keys
}
