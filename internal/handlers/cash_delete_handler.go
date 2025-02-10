package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
)

type CashDeleteHandler struct {
	cashService CashService
}

func NewCashDeleteHandler(cashService CashService) *CashDeleteHandler {
	return &CashDeleteHandler{
		cashService: cashService,
	}
}

func (h *CashDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.handle(r); err != nil {
		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte("deleted"))
}

func (h *CashDeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}
