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

type UpdateItemHandler struct {
	itemGetter       services.ItemGetter
	itemUpdater      services.ItemUpdater
	categoriesGetter services.CategoriesGetter
	defaultCurrency  *models.Currency
}

func NewUpdateItemHandler(app *app.App) http.Handler {
	itemRepository := repositories.NewItemRepository(app.DB, app.Currencies)

	return &UpdateItemHandler{
		itemGetter:  services.NewGetItemService(itemRepository),
		itemUpdater: services.NewUpdateItemService(itemRepository),
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
		defaultCurrency: app.DefaultCurrency,
	}
}

//nolint:funlen,cyclop
func (h *UpdateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currency, ok := r.Context().Value(models.CurrencyContextValue{}).(*models.Currency)
	if !ok {
		currency = h.defaultCurrency
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

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	_ = item.
		SetDate(r.FormValue("date")).
		SetFormula(r.FormValue("formula")).
		SetCategoryID(r.FormValue("category_id")).
		SetDescription(r.FormValue("description"))

	categories, err := h.categoriesGetter.GetCategories(r.Context(), currency.ID)
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if !item.IsValid() {
		if err := itemcomponents.EditPage(item, categories).Render(r.Context(), w); err != nil {
			slog.Error(err.Error())
		}

		return
	}

	if err := h.itemUpdater.UpdateItem(r.Context(), item.Item); err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	item, err = h.itemGetter.GetItem(r.Context(), r.PathValue("id"))
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := itemcomponents.UpdatePage(item).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
