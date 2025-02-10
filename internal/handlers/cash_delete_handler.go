package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/responses"
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
		if response, ok := w.(*responses.Response); ok {
			response.Error = err

			return
		}

		return
	}

	_, _ = w.Write([]byte("deleted"))
}

func (h *CashDeleteHandler) handle(r *http.Request) error {
	return h.cashService.Delete(r.Context(), r.PathValue("id"))
}
