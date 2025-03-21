package params

import (
	"maps"
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

func (p Params) With(key, value string) Params {
	values := maps.Clone(p.values)

	values.Set(key, value)

	return Params{values}
}

func (p Params) String() string {
	return p.values.Encode()
}

func (p Params) WithCurrency(cid currency.Currency) Params {
	return p.With("currency", currency.GetCode(cid))
}

func (p Params) WithMonth(month int) Params {
	return p.With("month", strconv.Itoa(month))
}

func (p Params) WithYear(year int) Params {
	return p.With("year", strconv.Itoa(year))
}

func (p Params) WithCategory(category int) Params {
	return p.With("category", strconv.Itoa(category))
}
