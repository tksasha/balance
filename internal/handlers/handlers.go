package handlers

import "github.com/tksasha/balance/internal/services"

type Handlers []Handler

func GetHandlers(
	itemService services.ItemService,
	categoryService services.CategoryService,
) []Handler {
	return []Handler{
		NewGetItemsHandler(itemService),
		NewCreateItemHandler(itemService),
		NewGetCategoriesHandler(categoryService),
		NewIndexPageHandler(),
	}
}
