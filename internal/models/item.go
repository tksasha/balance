package models

import (
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Item struct {
	ID           int
	Date         time.Time
	Sum          float32
	CategoryName string
	Description  string
}

func (i *Item) GetDate() string {
	return i.Date.Format(time.DateOnly)
}

func (i *Item) GetSum() string {
	printer := message.NewPrinter(language.Ukrainian)

	return printer.Sprintf("%.02f", i.Sum)
}
