package item

import (
	"database/sql"
	"time"

	"github.com/tksasha/balance/internal/common/currency"
)

type Item struct {
	ID           int
	Date         time.Time
	Formula      string
	Sum          float64
	CategoryID   int
	CategoryName sql.NullString
	CategorySlug string
	Description  string
	Currency     currency.Currency
}

type Items []*Item
