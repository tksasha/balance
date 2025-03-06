package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/index/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
)

type Handler struct {
	*commonhandler.Handler

	component *component.Component
}

func New() *Handler {
	return &Handler{
		Handler:   commonhandler.New(),
		component: component.New(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.component.Index(r.URL.Query()).Render(w)

	h.SetError(w, err)
}
