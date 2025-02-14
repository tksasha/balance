package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/category"
	"github.com/tksasha/balance/internal/core/category/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
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
		handlers.SetError(w, err)

		return
	}

	err = components.Edit(category).Render(w)

	handlers.SetError(w, err)
}

func (h *EditHandler) handle(r *http.Request) (*category.Category, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
