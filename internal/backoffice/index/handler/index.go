package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/index/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	commonhandler.Handler

	component *component.Component
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{
		component: component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Hx-Trigger-After-Swap", "backoffice.index.shown")

	err := h.component.Index().Render(w)

	h.SetError(w, err)
}
