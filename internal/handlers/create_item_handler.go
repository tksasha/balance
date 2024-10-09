package handlers

import (
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type CreateItemHandler struct {
	itemCreator      repositories.ItemCreator
	itemsGetter      services.ItemsGetter
	categoriesGetter services.CategoriesGetter
}

func NewCreateItemHandler(app *app.App) http.Handler {
	itemRepository := repositories.NewItemRepository(app.DB)

	return &CreateItemHandler{
		itemCreator: itemRepository,
		itemsGetter: services.NewGetItemsService(itemRepository),
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
	}
}

func (h *CreateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	item := models.NewItem().
		SetDate(r.FormValue("date")).
		SetFormula(r.FormValue("formula")).
		SetCategoryID(r.FormValue("category_id")).
		SetDescription(r.FormValue("description"))

	categories, err := h.categoriesGetter.GetCategories(r.Context(), 0) // TODO: use currency instead
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if !item.IsValid() {
		if err := itemcomponents.Form(decorators.NewItemDecorator(item), categories).Render(r.Context(), w); err != nil {
			slog.Error(err.Error())
		}

		return
	}

	// TODO: use CreateItemService.CreateItem()
	if err := services.CreateItemService(r.Context(), h.itemCreator, item); err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	items, err := h.itemsGetter.GetItems(r.Context())
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := itemcomponents.CreatePage(items, categories).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
