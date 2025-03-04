package handlers

import (
	"net/http"

	"github.com/tksasha/balance/internal/app/categoryreport"
	"github.com/tksasha/balance/internal/app/categoryreport/component"
	commonhandler "github.com/tksasha/balance/internal/common/handler"
)

type ShowHandler struct {
	*commonhandler.Handler

	service   categoryreport.Service
	component *component.Component
}

func NewShowHandler(service categoryreport.Service) *ShowHandler {
	return &ShowHandler{
		Handler:   commonhandler.New(),
		service:   service,
		component: component.New(),
	}
}

func (h *ShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := categoryreport.Request{
		Year:  r.URL.Query().Get("year"),
		Month: r.URL.Query().Get("month"),
	}

	entities, err := h.service.Report(r.Context(), request)
	if err != nil {
		h.SetError(w, err)
	}

	err = h.component.Show(entities).Render(w)

	h.SetError(w, err)
}
