package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/models"
	"github.com/tksasha/balance/internal/responses"
)

type CashListHandler struct {
	cashService CashService
}

func NewCashListHandler(cashService CashService) *CashListHandler {
	return &CashListHandler{
		cashService: cashService,
	}
}

func (h *CashListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		if response, ok := w.(*responses.Response); ok {
			response.Error = err
		}

		return
	}

	_ = cashes

	_, _ = w.Write([]byte(""))
}

func (h *CashListHandler) handle(r *http.Request) (models.Cashes, error) {
	return h.cashService.List(r.Context())
}
