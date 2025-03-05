package balance

import (
	"github.com/shopspring/decimal"
)

type Balance struct {
	AtEnd   decimal.Decimal
	Balance decimal.Decimal
}
