package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/components"
	"github.com/tksasha/balance/internal/common"
)

type NewHandler struct {
	*common.BaseHandler

	cashComponent *components.CashComponent
}

func NewNewHandler(
	baseHandler *common.BaseHandler,
	cashComponent *components.CashComponent,
) *NewHandler {
	return &NewHandler{
		BaseHandler:   baseHandler,
		cashComponent: cashComponent,
	}
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cash := &cash.Cash{}

	err := h.cashComponent.New(cash).Render(w)

	h.SetError(w, err)
}
