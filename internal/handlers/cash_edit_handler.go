package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/responses"
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
		if rw, ok := w.(*responses.Response); ok {
			rw.Error = err

			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	_, _ = w.Write([]byte(cash.Name))
}

func (h *CashEditHandler) handle(r *http.Request) (*models.Cash, error) {
	return h.cashService.FindByID(r.Context(), r.PathValue("id"))
}
