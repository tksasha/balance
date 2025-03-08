package currency

import "strings"

const (
	UAH Currency = iota + 1
	USD
	EUR

	Default = UAH
)

type Currency int

type ContextValue struct{}

func GetByCode(code string) Currency {
	return map[string]Currency{
		"uah": UAH,
		"usd": USD,
		"eur": EUR,
	}[strings.ToLower(code)]
}

func GetCode(currency Currency) string {
	code, ok := map[Currency]string{
		UAH: "uah",
		USD: "usd",
		EUR: "eur",
	}[currency]
	if ok {
		return code
	}

	return GetCode(Default)
}
