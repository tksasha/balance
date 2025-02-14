package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/cash/components"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type EditHandler struct {
	service cash.Service
}

func NewEditHandler(service cash.Service) *EditHandler {
	return &EditHandler{
		service: service,
	}
}

func (h *EditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		handlers.E(w, err)

		return
	}

	if err := components.Edit(cash).Render(w); err != nil {
		handlers.E(w, err)
	}
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
