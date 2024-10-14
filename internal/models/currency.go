package models

type Currency struct {
	ID   int
	Code string
}

type CurrencyContextValue struct{}

func NewDefaultCurrency() *Currency {
	return &Currency{ID: 0, Code: "UAH"}
}
