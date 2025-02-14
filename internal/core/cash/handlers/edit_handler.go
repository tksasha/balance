package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/common/handlers"
	"github.com/tksasha/balance/internal/core/cash"
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

	_, _ = w.Write([]byte(cash.Name))
}

func (h *EditHandler) handle(r *http.Request) (*cash.Cash, error) {
	return h.service.FindByID(r.Context(), r.PathValue("id"))
}
