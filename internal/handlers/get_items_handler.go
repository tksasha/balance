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
	currency    models.Currency
	itemsGetter services.ItemsGetter
}

func NewGetItemsHandler(currency models.Currency, app *app.App) http.Handler {
	return &GetItemsHandler{
		currency: currency,
		itemsGetter: services.NewGetItemsService(
			repositories.NewItemRepository(app.DB),
		),
	}
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.itemsGetter.GetItems(r.Context(), h.currency)
	if err != nil {
		slog.Error(err.Error())
	}

	if err := itemcomponents.IndexPage(h.currency, items).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
