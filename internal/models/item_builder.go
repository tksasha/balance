package models

import (
	"strconv"
	"strings"
	"time"

	"github.com/tksasha/calculator"
	"github.com/tksasha/model"
	"github.com/tksasha/model/errors"
	"github.com/tksasha/model/errors/messages"
)

type ItemBuilder struct {
	Model *model.Model
	item  *Item
}

func NewItemBuilder() *ItemBuilder {
	return &ItemBuilder{
		Model: model.New(),
		item:  &Item{},
	}
}

func (b *ItemBuilder) WithDate(input string) *ItemBuilder {
	if input == "" {
		b.Model.Errors.Set("date", messages.Required)
	}

	date, err := time.Parse(time.DateOnly, input)
	if err != nil {
		b.Model.Errors.Set("date", messages.Invalid)
	}

	b.item.Date = date

	return b
}

func (b *ItemBuilder) WithFormula(formula string) *ItemBuilder {
	if formula == "" {
		b.Model.Errors.Set("formula", messages.Required)
	}

	sum, err := calculator.Calculate(formula)
	if err != nil {
		b.Model.Errors.Set("formula", messages.Invalid)
	}

	b.item.Formula = formula
	b.item.Sum = sum

	return b
}

func (b *ItemBuilder) WithCategoryID(input string) *ItemBuilder {
	if input == "" {
		b.Model.Errors.Set("category_id", messages.Required)
	}

	categoryID, err := strconv.Atoi(input)
	if err != nil {
		b.Model.Errors.Set("category_id", messages.Invalid)
	}

	b.item.CategoryID = categoryID

	return b
}

func (b *ItemBuilder) WithDescription(input string) *ItemBuilder {
	b.item.Description = strings.TrimSpace(input)

	return b
}

func (b *ItemBuilder) WithCurrencyCode(code string) *ItemBuilder {
	currency := GetCurrencyByCode(code)
	if currency == 0 {
		currency, code = GetDefaultCurrency()
	}

	b.item.Currency = currency
	b.item.CurrencyCode = code

	return b
}

func (b *ItemBuilder) Build() (*Item, errors.ValidationError) {
	if !b.Model.Errors.IsEmpty() {
		return nil, b.Model.Errors
	}

	return b.item, nil
}
