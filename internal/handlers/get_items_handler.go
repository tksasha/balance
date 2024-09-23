package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/tksasha/balance/internal/decorators"
	"github.com/tksasha/balance/internal/repositories"
	"github.com/tksasha/balance/internal/server/app"
)

type GetItemsHandler struct {
	template       *template.Template
	itemRepository *repositories.ItemRepository
}

func NewGetItemsHandler(app *app.App) http.Handler {
	return &GetItemsHandler{
		template:       app.T,
		itemRepository: repositories.NewItemRepository(app.DB),
	}
}

func (h *GetItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.itemRepository.GetItems(r.Context())
	if err != nil {
		slog.Error(err.Error())
	}

	decoratedItems := make([]*decorators.ItemDecorator, 0, len(items))

	for _, item := range items {
		decoratedItems = append(decoratedItems, decorators.NewItemDecorator(&item))
	}

	if err := h.template.ExecuteTemplate(w, "item-list", decoratedItems); err != nil {
		slog.Error(err.Error())
	}
}
