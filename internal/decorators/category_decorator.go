package decorators

import (
	"strconv"

	"github.com/tksasha/balance/internal/models"
)

type CategoryDecorator struct {
	*models.Category
	ID string
}

func NewCategoryDecorator(category *models.Category) *CategoryDecorator {
	return &CategoryDecorator{
		Category: category,
		ID:       strconv.Itoa(category.ID),
	}
}
