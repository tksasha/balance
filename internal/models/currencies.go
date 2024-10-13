package models

type Currencies map[string]Currency

func (currencies Currencies) GetByID(id int) *Currency {
	for _, currency := range currencies {
		if currency.ID == id {
			return &currency
		}
	}

	return nil
}
