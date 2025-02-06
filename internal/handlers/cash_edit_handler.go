package handlers

import (
	"errors"
	"net/http"

	"github.com/tksasha/balance/internal/apperrors"
	"github.com/tksasha/balance/internal/models"
)

type CashEditHandler struct {
	cashService CashService
}

func NewCashEditHandler(cashService CashService) *CashEditHandler {
	return &CashEditHandler{
		cashService: cashService,
	}
}

func (h *CashEditHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash, err := h.handle(r)
	if err != nil {
		if errors.Is(err, apperrors.ErrResourceNotFound) {
			http.Error(w, "Resource Not Found", http.StatusNotFound)

			return
		}
	}

	_, _ = w.Write([]byte(cash.Name))
}

func (h *CashEditHandler) handle(r *http.Request) (*models.Cash, error) {
	return h.cashService.FindByID(r.Context(), r.PathValue("id"))
}
