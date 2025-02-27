package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/backoffice/category"
	"github.com/tksasha/balance/internal/backoffice/category/component"
	"github.com/tksasha/balance/internal/common/handler"
)

type ListHandler struct {
	*handler.Handler

	service   category.Service
	component *component.Component
}

func NewListHandler(
	service category.Service,
) *ListHandler {
	return &ListHandler{
		Handler:   handler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.List(categories).Render(w)

	h.SetError(w, err)
}

func (h *ListHandler) handle(r *http.Request) (category.Categories, error) {
	return h.service.List(r.Context())
}
