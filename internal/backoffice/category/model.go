package category

import (
	"database/sql"

	"github.com/tksasha/balance/internal/common/currency"
)

type Category struct {
	ID            int
	Name          string
	Slug          string
	Income        bool
	Visible       bool
	Currency      currency.Currency
	Supercategory int
	DeletedAt     sql.NullTime
}

type Categories []*Category
