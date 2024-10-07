package handlers

import (
	"log/slog"
	"net/http"

	itemcomponents "github.com/tksasha/balance/internal/components/items"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
	"github.com/tksasha/balance/internal/services"
)

type UpdateItemHandler struct {
	itemUpdater services.ItemUpdater
}

func NewUpdateItemHandler(app *app.App) http.Handler {
	itemRepository := repositories.NewItemRepository(app.DB)

	itemGetter := services.NewGetItemService(itemRepository)

	return &UpdateItemHandler{
		services.NewUpdateItemService(
			itemGetter,
			itemRepository,
		),
	}
}

func (h *UpdateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if err := itemcomponents.UpdatePage().Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
