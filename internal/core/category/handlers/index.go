package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type IndexHandler struct {
	service category.Service
}

func NewIndexHandler(service category.Service) *IndexHandler {
	return &IndexHandler{
		service: service,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		handlers.SetError(w, err)

		return
	}

	err = components.Index(categories).Render(w)

	handlers.SetError(w, err)
}

func (h *IndexHandler) handle(r *http.Request) (category.Categories, error) {
	return h.service.List(r.Context())
}
