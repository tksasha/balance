package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/components"
	"github.com/tksasha/balance/internal/models"
)

type CashNewHandler struct{}

func NewCashNewHandler() *CashNewHandler {
	return &CashNewHandler{}
}

func (h *CashNewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &models.Cash{}

	if err := components.CashNew(cash).Render(w); err != nil {
		e(w, err)
	}
}
