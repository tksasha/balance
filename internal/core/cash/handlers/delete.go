package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/core/cash"
	"github.com/tksasha/balance/internal/core/common/handlers"
)

type DeleteHandler struct {
	cashService cash.Service
}

func NewDeleteHandler(cashService cash.Service) *DeleteHandler {
	return &DeleteHandler{
		cashService: cashService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		handlers.SetError(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}
