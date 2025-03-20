package currency

import "strings"

const (
	UAH Currency = iota + 1
	USD
	EUR

	Default = UAH
)

type Currency int

type Currencies map[Currency]string

type ContextValue struct{}

func All() Currencies {
	return Currencies{
		UAH: "UAH",
		USD: "USD",
		EUR: "EUR",
	}
}

func GetByCode(code string) Currency {
	for curr, currCode := range All() {
		if Code(code) == Code(currCode) {
			return curr
		}
	}

	return Default
}

func GetCode(currency Currency) string {
	code, ok := All()[currency]
	if ok {
		return Code(code)
	}

	return GetCode(Default)
}

func Code(code string) string {
	return strings.ToLower(code)
}
