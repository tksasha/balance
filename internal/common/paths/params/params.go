package params

import (
	"net/url"
	"strconv"

	"github.com/tksasha/balance/internal/common/currency"
)

type Params struct {
	values url.Values
}

func New(args ...url.Values) Params {
	values := url.Values{
		"currency": {currency.GetCode(currency.Default)},
	}

	if len(args) == 1 {
		values = args[0]
	}

	params := Params{
		values: values,
	}

	return params
}

func (p Params) Get(key string) string {
	return p.values.Get(key)
}

func (p Params) Has(key string) bool {
	return p.values.Has(key)
}

func (p Params) Set(key, value string) {
	p.values.Set(key, value)
}

func (p Params) String() string {
	return p.values.Encode()
}

func (p Params) SetCurrency(cid currency.Currency) Params {
	p.values.Set("currency", currency.GetCode(cid))

	return p
}

func (p Params) SetCurrencyCode(code string) Params {
	p.SetCurrency(currency.GetByCode(code))

	return p
}

func (p Params) SetMonth(month int) Params {
	p.values.Set("month", strconv.Itoa(month))

	return p
}

func (p Params) SetYear(year int) Params {
	p.values.Set("year", strconv.Itoa(year))

	return p
}

func (p Params) SetCategory(category int) Params {
	p.values.Set("category", strconv.Itoa(category))

	return p
}
