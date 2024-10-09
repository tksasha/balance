package decorators

import "github.com/tksasha/balance/internal/models"

type CategoriesDecorator struct {
	Income  []*CategoryDecorator
	Expense []*CategoryDecorator
}

func NewCategoriesDecorator(categories []*models.Category) *CategoriesDecorator {
	decorator := new(CategoriesDecorator)

	for _, category := range categories {
		if category.Income {
			decorator.Income = append(decorator.Income, NewCategoryDecorator(category))

			continue
		}

		decorator.Expense = append(decorator.Expense, NewCategoryDecorator(category))
	}

	return decorator
}
