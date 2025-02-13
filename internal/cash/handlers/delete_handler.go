package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/cash"
	"github.com/tksasha/balance/internal/common/handlers"
)

type DeleteHandler struct {
	service cash.Service
}

func NewDeleteHandler(service cash.Service) *DeleteHandler {
	return &DeleteHandler{
		service: service,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		handlers.E(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.service.Delete(r.Context(), r.PathValue("id"))
}
