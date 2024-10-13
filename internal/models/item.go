package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/tksasha/calculator"
	"github.com/tksasha/model"
	"github.com/tksasha/model/errors/messages"
)

type Item struct {
	*model.Model
	ID           int
	Date         time.Time
	Sum          float64
	Formula      string
	CategoryID   int
	CategoryName string
	Description  string
	Currency     Currency
	currencies   Currencies
	CurrencyID   int
}

func NewItem(currencies Currencies) *Item {
	item := new(Item)

	item.Model = model.New()

	item.currencies = currencies

	return item
}

func (i *Item) SetDate(value string) *Item {
	if value == "" {
		i.Errors.Set("date", messages.Required)
	}

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		i.Errors.Set("date", messages.Invalid)
	}

	i.Date = date

	return i
}

func (i *Item) SetFormula(formula string) *Item {
	if formula == "" {
		i.Errors.Set("formula", messages.Required)
	}

	sum, err := calculator.Calculate(formula)
	if err != nil {
		i.Errors.Set("formula", messages.Invalid)
	}

	i.Formula = formula

	i.Sum = sum

	return i
}

func (i *Item) SetCategoryID(value string) *Item {
	if value == "" {
		i.Errors.Set("category_id", messages.Required)
	}

	categoryID, err := strconv.Atoi(value)
	if err != nil {
		i.Errors.Set("category_id", messages.Invalid)
	}

	i.CategoryID = categoryID

	return i
}

func (i *Item) SetDescription(description string) *Item {
	i.Description = description

	return i
}

func (i *Item) SetCurrency(code string) *Item {
	code = strings.ToUpper(code)

	if code == "" {
		i.Errors.Set("currency", messages.Required)
	}

	currency, ok := i.currencies[code]
	if !ok {
		i.Errors.Set("currency", messages.Invalid)
	}

	i.Currency = currency

	return i
}
