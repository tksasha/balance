package handlers

import (
	"log/slog"
	"net/http"
	"time"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
)

type GetItemsHandler struct {
	itemRepository *repositories.ItemRepository
}

func NewGetItemsHandler(app *app.App) http.Handler {
	return &GetItemsHandler{
		itemRepository: repositories.NewItemRepository(app.DB),
	}
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)

	items, err := h.itemRepository.GetItems(r.Context())
	if err != nil {
		slog.Error(err.Error())
	}

	if err := itemcomponents.Table(items).Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
