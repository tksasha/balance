package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/common/handlers"
	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
)

type EditHandler struct {
	service category.Service
}

func NewEditHandler(service category.Service) *EditHandler {
	return &EditHandler{
		service: service,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	category, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Edit(category).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *EditHandler) handle(r *http.Request) (*category.Category, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
