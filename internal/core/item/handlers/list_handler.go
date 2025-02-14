package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type ListHandler struct {
	service item.Service
}

func NewListHandler(service item.Service) *ListHandler {
	return &ListHandler{
		service: service,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Index(items).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *ListHandler) handle(r *http.Request) (item.Items, error) {
	return h.service.List(r.Context())
}
