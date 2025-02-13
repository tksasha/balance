package category

import (
	"database/sql"

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
