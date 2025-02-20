package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type ListHandler struct {
	*handlers.BaseHandler

	itemService    item.Service
	itemsComponent *components.ItemsComponent
}

func NewListHandler(
	baseHandler *handlers.BaseHandler,
	itemService item.Service,
	itemsComponent *components.ItemsComponent,
) *ListHandler {
	return &ListHandler{
		BaseHandler:    baseHandler,
		itemService:    itemService,
		itemsComponent: itemsComponent,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.itemsComponent.Index(items).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (item.Items, error) {
	request := item.IndexRequest{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	return h.itemService.Index(r.Context(), request)
}
