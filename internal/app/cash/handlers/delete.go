package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/common/handler"
)

type DeleteHandler struct {
	*handler.Handler

	cashService cash.Service
}

func NewDeleteHandler(cashService cash.Service) *DeleteHandler {
	return &DeleteHandler{
		Handler:     handler.New(),
		cashService: cashService,
	}
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		h.SetError(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}
