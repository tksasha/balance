package path

import "github.com/tksasha/balance/internal/common/currency"

type Params map[string]string

func NewCurrency(curr currency.Currency) Params {
	return Params{
		"currency": currency.GetCode(curr),
	}
}
