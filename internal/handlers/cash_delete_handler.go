package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/handlers/utils"
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
		utils.E(w, err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CashDeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}
