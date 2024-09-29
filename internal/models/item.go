package models

import (
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Item struct {
	ID           int
	Date         time.Time
	Sum          float64
	Formula      string
	CategoryID   int
	CategoryName string
	Description  string
}

func NewItem(
	date time.Time,
	formula string,
	sum float64,
	categoryID int,
	categoryName string,
	description string,
) *Item {
	return &Item{
		Date:         date,
		Formula:      formula,
		Sum:          sum,
		CategoryID:   categoryID,
		CategoryName: categoryName,
		Description:  description,
	}
}

func (i *Item) GetDate() string {
	return i.Date.Format(time.DateOnly)
}

func (i *Item) GetSum() string {
	printer := message.NewPrinter(language.Ukrainian)

	return printer.Sprintf("%.02f", i.Sum)
}
