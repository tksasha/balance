package handlers

import (
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
	currency         models.Currency
	currencies       models.Currencies
}

func NewIndexHandler(app *app.App) http.Handler {
	return &IndexHandler{
		categoriesGetter: services.NewGetCategoriesService(
			repositories.NewCategoryRepository(app.DB),
		),
		currency:   app.Currency,
		currencies: app.Currencies,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currency, ok := r.Context().Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency = h.currency
	}

	categories, err := h.categoriesGetter.GetCategories(r.Context(), currency.ID)
	if err != nil {
		panic(err)
	}

	itemDecorator := decorators.NewItemDecorator(
		models.NewItem(h.currencies).SetCurrency(currency.Code),
	)

	if err := indexcomponents.IndexPage(currency, categories, itemDecorator).Render(r.Context(), w); err != nil {
		panic(err)
	}
}
