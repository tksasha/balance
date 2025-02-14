package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/common/handlers"
	"github.com/tksasha/balance/internal/core/item"
	"github.com/tksasha/balance/internal/core/item/components"
)

type IndexHandler struct {
	service item.Service
}

func NewIndexHandler(service item.Service) *IndexHandler {
	return &IndexHandler{
		service: service,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Index(items).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *IndexHandler) handle(r *http.Request) (item.Items, error) {
	return h.service.List(r.Context())
}
