package decorators

import (
	"time"

	"github.com/tksasha/balance/internal/models"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type ItemDecorator struct {
	*models.Item

	Date string
	Sum  string
}

func NewItemDecorator(item *models.Item) *ItemDecorator {
	decorator := &ItemDecorator{
		Item: item,
		Sum:  message.NewPrinter(language.Ukrainian).Sprintf("%0.2f", item.Sum),
	}

	if !item.Date.IsZero() {
		decorator.Date = item.Date.Format(time.DateOnly)
	}

	return decorator
}
