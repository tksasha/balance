package handlers

import (
	"fmt"
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
	categoriesGetter services.CategoriesGetter
	defaultCurrency  *models.Currency
	currencies       models.Currencies
}

func NewIndexHandler(app *app.App) http.Handler {
	return &IndexHandler{
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
		defaultCurrency: app.DefaultCurrency,
		currencies:      app.Currencies,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(w, r); err != nil {
		slog.Error("index handler error", "error", err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *IndexHandler) handle(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	currency, ok := ctx.Value(models.CurrencyContextValue{}).(*models.Currency)
	if !ok {
		currency = h.defaultCurrency
	}

	categories, err := h.categoriesGetter.GetCategories(ctx, currency.ID)
	if err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}

	itemDecorator := decorators.NewItemDecorator(
		models.NewItem(h.currencies).SetCurrency(currency.Code),
	)

	if err := indexcomponents.IndexPage(currency, categories, itemDecorator).Render(ctx, w); err != nil {
		return fmt.Errorf("failed to render index page: %w", err)
	}

	return nil
}
