package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	internalerrors "github.com/tksasha/balance/internal/errors"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type EditItemHandler struct {
	itemGetter       services.ItemGetter
	categoriesGetter services.CategoriesGetter
	currency         models.Currency
}

func NewEditItemHandler(app *app.App) http.Handler {
	return &EditItemHandler{
		itemGetter: services.NewGetItemService(
			repositories.NewItemRepository(app.DB, app.Currencies),
		),
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
		currency: app.Currency,
	}
}

func (h *EditItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currency, ok := r.Context().Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency = h.currency
	}

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

	categories, err := h.categoriesGetter.GetCategories(r.Context(), currency.ID)
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := itemcomponents.EditPage(item, categories).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
