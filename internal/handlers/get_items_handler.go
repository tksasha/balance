package handlers

import (
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type GetItemsHandler struct {
	itemsGetter services.ItemsGetter
	currency    models.Currency
}

func NewGetItemsHandler(app *app.App) http.Handler {
	return &GetItemsHandler{
		itemsGetter: services.NewGetItemsService(
			repositories.NewItemRepository(app.DB, app.Currencies),
		),
		currency: app.Currency,
	}
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currency, ok := r.Context().Value(models.CurrencyContextValue{}).(models.Currency)
	if !ok {
		currency = h.currency
	}

	items, err := h.itemsGetter.GetItems(r.Context(), currency)
	if err != nil {
		slog.Error(err.Error())
	}

	if err := itemcomponents.IndexPage(items).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
