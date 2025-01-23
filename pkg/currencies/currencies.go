package currencies

import "strings"

const (
	UAH Currency = iota + 1
	USD
	EUR

	DefaultCurrency = UAH
)

type Currency int

type CurrencyContextValue struct{}

func GetCurrencyByCode(code string) Currency {
	return map[string]Currency{
		"uah": UAH,
		"usd": USD,
		"eur": EUR,
	}[strings.ToLower(code)]
}
