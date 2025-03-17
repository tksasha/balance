package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/cash"
	"github.com/tksasha/balance/internal/app/cash/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
)

type IndexHandler struct {
	*commonhandler.Handler

	service   cash.Service
	component *component.Component
}

func NewIndexHandler(service cash.Service) *IndexHandler {
	return &IndexHandler{
		Handler:   commonhandler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cashes, err := h.handle(r)
	if err != nil {
		h.SetError(w, err)

		return
	}

	err = h.component.Index(r.URL.Query(), cashes).Render(w)

	h.SetError(w, err)
}

func (h *IndexHandler) handle(r *http.Request) (cash.GroupedCashes, error) {
	return h.service.GroupedList(r.Context())
}
