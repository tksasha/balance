package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type EditItemHandler struct {
	itemGetter       services.ItemGetter
	categoriesGetter services.CategoriesGetter
}

func NewEditItemHandler(
	app *app.App,
) http.Handler {
	return &EditItemHandler{
		itemGetter: services.NewGetItemService(
			repositories.NewItemRepository(app.DB),
		),
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
	}
}

func (h *EditItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	item, err := h.itemGetter.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		if errors.Is(err, internalerrors.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	categories, err := h.categoriesGetter.GetCategories(r.Context(), 0)
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := itemcomponents.EditPage(item, categories).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
