package models

import "strings"

const (
	UAH Currency = iota + 1
	USD
	EUR
)

var currencyCodeToID = map[string]Currency{ //nolint:gochecknoglobals
	"uah": UAH,
	"usd": USD,
	"eur": EUR,
}

type Currency int

type CurrencyContextValue struct{}

func GetCurrencyByCode(code string) Currency {
	return currencyCodeToID[strings.ToLower(code)]
}

func GetDefaultCurrency() (Currency, string) {
	return UAH, "uah"
}
