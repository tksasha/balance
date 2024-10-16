package models

import (
	"strconv"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Item struct {
	ID           int
	Date         time.Time
	Formula      string
	Sum          float64
	CategoryID   int
	CategoryName string
	Description  string
	Currency     Currency
	CurrencyCode string
}

type Items []*Item

func (i *Item) GetIDAsString() string {
	return strconv.Itoa(i.ID)
}

func (i *Item) GetDateAsString() string {
	if i.Date.IsZero() {
		return ""
	}

	return i.Date.Format(time.DateOnly)
}

func (i *Item) GetSumAsString() string {
	return message.NewPrinter(language.Ukrainian).Sprintf("%0.2f", i.Sum)
}
