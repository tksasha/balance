package handlers

import (
	"log/slog"
	"net/http"

	indexcomponents "github.com/tksasha/balance/internal/components/index"
	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type IndexHandler struct {
	currency         models.Currency
	categoriesGetter services.CategoriesGetter
}

func NewIndexHandler(currency models.Currency, app *app.App) http.Handler {
	return &IndexHandler{
		currency: currency,
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoriesGetter.GetCategories(r.Context(), h.currency.ID)
	if err != nil {
		slog.Error(err.Error())

		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	itemDecorator := decorators.NewItemDecorator(models.NewItem())

	if err := indexcomponents.IndexPage(h.currency, categories, itemDecorator).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
