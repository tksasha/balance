package handlers

import (
	"net/http"

	components "github.com/tksasha/balance/internal/components/cash"
	"github.com/tksasha/balance/internal/models"
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
		e(w, err)

		return
	}

	if err := components.CashList(cashes).Render(w); err != nil {
		e(w, err)
	}
}

func (h *CashListHandler) handle(r *http.Request) (models.Cashes, error) {
	return h.cashService.List(r.Context())
}
