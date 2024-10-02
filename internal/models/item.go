package models

import (
	"strconv"
	"time"

	"github.com/tksasha/calculator"
	"github.com/tksasha/model"
	"github.com/tksasha/model/errors/messages"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
}

func NewItem() *Item {
	item := new(Item)

	item.Model = model.New()

	return item
}

func (i *Item) SetDate(value string) {
	if value == "" {
		i.Errors.Set("date", messages.Required)
	}

	date, err := time.Parse(time.DateOnly, value)
	if err != nil {
		i.Errors.Set("date", messages.Invalid)
	}

	i.Date = date
}

func (i *Item) GetDateAsString() string {
	return i.Date.Format(time.DateOnly)
}

func (i *Item) SetFormula(formula string) {
	if formula == "" {
		i.Errors.Set("formula", messages.Required)
	}

	sum, err := calculator.Calculate(formula)
	if err != nil {
		i.Errors.Set("formula", messages.Invalid)
	}

	i.Formula = formula

	i.Sum = sum
}

func (i *Item) GetSumAsString() string {
	printer := message.NewPrinter(language.Ukrainian)

	return printer.Sprintf("%0.2f", i.Sum)
}

func (i *Item) SetCategoryID(value string) {
	if value == "" {
		i.Errors.Set("category_id", messages.Required)
	}

	categoryID, err := strconv.Atoi(value)
	if err != nil {
		i.Errors.Set("category_id", messages.Invalid)
	}

	i.CategoryID = categoryID
}

func (i *Item) GetCategoryIDAsString() string {
	return strconv.Itoa(i.CategoryID)
}

func (i *Item) SetDescription(description string) {
	i.Description = description
}

func (i *Item) GetDescription() string {
	return i.Description
}
