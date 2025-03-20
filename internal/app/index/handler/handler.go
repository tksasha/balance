package handler

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/index/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
	"github.com/tksasha/balance/internal/common/paths/params"
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
	params := params.New(r.URL.Query())

	err := h.component.Index(params).Render(w)

	h.SetError(w, err)
}
