package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/balance"
	"github.com/tksasha/balance/internal/app/balance/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
)

type ShowHandler struct {
	*commonhandler.Handler

	service   balance.Service
	component *component.Component
}

func NewShowHandler(service balance.Service) *ShowHandler {
	return &ShowHandler{
		Handler:   commonhandler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *ShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	balance, err := h.service.Balance(r.Context())
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.Balance(balance).Render(w)

	h.SetError(w, err)
}
