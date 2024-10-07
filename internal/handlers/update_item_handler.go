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
	getItemService    *services.GetItemService
	updateItemService *services.UpdateItemService
}

func NewUpdateItemHandler(app *app.App) http.Handler {
	itemRepository := repositories.NewItemRepository(app.DB)

	return &UpdateItemHandler{
		getItemService:    services.NewGetItemService(itemRepository),
		updateItemService: services.NewUpdateItemService(itemRepository),
	}
}

func (h *UpdateItemHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := itemcomponents.UpdatePage().Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
	}
}
