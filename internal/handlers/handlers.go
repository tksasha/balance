package handlers

type Handlers []Handler

func New(
	itemService ItemService,
	categoryService CategoryService,
) Handlers {
	return Handlers{
		NewGetItemsHandler(itemService),
		NewCreateItemHandler(itemService, categoryService),
		NewGetCategoriesHandler(categoryService),
		NewIndexPageHandler(categoryService),
		NewGetItemHandler(itemService),
	}
}
