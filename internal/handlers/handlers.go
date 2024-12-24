package handlers

type Handlers []Handler

func GetHandlers(
	itemService ItemService,
	categoryService CategoryService,
) []Handler {
	return []Handler{
		NewGetItemsHandler(itemService),
		NewCreateItemHandler(itemService, categoryService),
		NewGetCategoriesHandler(categoryService),
		NewIndexPageHandler(categoryService),
	}
}
