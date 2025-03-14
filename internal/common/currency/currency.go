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
		if strings.ToUpper(code) == currCode {
			return curr
		}
	}

	return Default
}

func GetCode(currency Currency) string {
	code, ok := All()[currency]
	if ok {
		return strings.ToLower(code)
	}

	return GetCode(Default)
}
