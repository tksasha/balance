package models

type Currencies map[string]*Currency

func NewCurrencies() Currencies {
	return Currencies{
		"UAH": {
			ID:   0,
			Code: "UAH",
		},
		"USD": {
			ID:   1,
			Code: "USD",
		},
		"EUR": {
			ID:   3, //nolint:mnd
			Code: "EUR",
		},
	}
}

func (currencies Currencies) GetByID(id int) *Currency {
	for _, currency := range currencies {
		if currency.ID == id {
			return currency
		}
	}

	return nil
}
