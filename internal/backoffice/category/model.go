package category

import "github.com/tksasha/balance/internal/common/currency"

type Category struct {
	ID       int
	Currency currency.Currency
	Name     string
}

type Categories []*Category
